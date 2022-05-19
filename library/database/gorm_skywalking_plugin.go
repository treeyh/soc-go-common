package database

/// gorm skywalking 插件，参考 https://github.com/go-gorm/opentracing 实现
import (
	"context"
	"github.com/SkyAPM/go2sky"
	jsoniter "github.com/json-iterator/go"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/library/tracing"
	"gorm.io/gorm"
	v3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// operationName defines a type to wrap the name of each operation name.
type operationName string

// String returns the actual string of operationName.
func (op operationName) String() string {
	return string(op)
}

const (
	_createOp operationName = "create"
	_updateOp operationName = "update"
	_queryOp  operationName = "query"
	_deleteOp operationName = "delete"
	_rowOp    operationName = "row"
	_rawOp    operationName = "raw"
)

// operationStage indicates the timing when the operation happens.
type operationStage string

// Name returns the actual string of operationStage.
func (op operationStage) Name() string {
	return string(op)
}

const (
	_stageBeforeCreate operationStage = "sky_walking:before_create"
	_stageAfterCreate  operationStage = "sky_walking:after_create"
	_stageBeforeUpdate operationStage = "sky_walking:before_update"
	_stageAfterUpdate  operationStage = "sky_walking:after_update"
	_stageBeforeQuery  operationStage = "sky_walking:before_query"
	_stageAfterQuery   operationStage = "sky_walking:after_query"
	_stageBeforeDelete operationStage = "sky_walking:before_delete"
	_stageAfterDelete  operationStage = "sky_walking:after_delete"
	_stageBeforeRow    operationStage = "sky_walking:before_row"
	_stageAfterRow     operationStage = "sky_walking:after_row"
	_stageBeforeRaw    operationStage = "sky_walking:before_raw"
	_stageAfterRaw     operationStage = "sky_walking:after_raw"
)

const (
	_prefix      = "gorm.sky_walking_tracing"
	_errorTagKey = "error"
)

var (
	// span.Tag keys
	_tableTagKey = keyWithPrefix("table")
	// span.Log keys
	//_errorLogKey        = keyWithPrefix("error")
	_resultLogKey       = keyWithPrefix("result")
	_sqlLogKey          = keyWithPrefix("sql")
	_rowsAffectedLogKey = keyWithPrefix("rowsAffected")
)

func keyWithPrefix(key string) string {
	return _prefix + "." + key
}

var (
	skyWalkingSpanKey = "sky_walking:span"
	json              = jsoniter.ConfigCompatibleWithStandardLibrary
)

type skyWalkingPlugin struct {
	// opt includes options those opentracingPlugin support.
	opt *options
}

func (p skyWalkingPlugin) injectBefore(db *gorm.DB, op operationName) {
	// make sure context could be used
	if db == nil || p.opt.tracer == nil {
		return
	}

	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "could not inject sp from nil Statement.Context or nil Statement")
		return
	}

	dbSpan, err := p.opt.tracer.CreateExitSpan(db.Statement.Context, op.String(), p.opt.url, func(key, value string) error {
		return nil
	})
	if err != nil {
		db.Logger.Error(context.TODO(), errors.SkyWalkingSpanNotInit.Error()+"; url:"+p.opt.url+"; op:"+op.String()+"; err:"+err.Error())
		return
	}
	dbSpan.SetComponent(tracing.GormComponent)
	dbSpan.SetSpanLayer(v3.SpanLayer_Database)

	db.InstanceSet(skyWalkingSpanKey, dbSpan)
}

func (p skyWalkingPlugin) extractAfter(db *gorm.DB) {
	// make sure context could be used
	if db == nil || p.opt.tracer == nil {
		return
	}
	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "could not extract sp from nil Statement.Context or nil Statement")
		return
	}

	// extract sp from db context
	//sp := opentracing.SpanFromContext(db.Statement.Context)
	v, ok := db.InstanceGet(skyWalkingSpanKey)
	if !ok || v == nil {
		return
	}

	sp, ok := v.(go2sky.Span)
	if !ok || sp == nil {
		return
	}
	defer sp.End()

	// tag and log fields we want.
	tag(&sp, db, p.opt.errorTagHook)
	sklog(&sp, db, p.opt.logResult, p.opt.logSqlParameters)
}

type options struct {
	// logResult means log SQL operation result into span log which causes span size grows up.
	// This is advised to only open in developing environment.
	logResult bool

	// tracer allows users to use customized and different tracer to makes tracing clearly.
	tracer *go2sky.Tracer

	url string

	// Whether to log statement parameters or leave placeholders in the queries.
	logSqlParameters bool

	// errorTagHook allows users to customized error what kind of error tag should be tagged.
	errorTagHook errorTagHook
}

func defaultOption() *options {
	return &options{
		logResult:        false,
		tracer:           nil,
		logSqlParameters: true,
		url:              "mysql",
		errorTagHook:     defaultErrorTagHook,
	}
}

type applyOption func(o *options)

// WithLogResult enable opentracingPlugin to log the result of each executed sql.
func WithLogResult(logResult bool) applyOption {
	return func(o *options) {
		o.logResult = logResult
	}
}

// WithTracer allows to use customized tracer rather than the global one only.
func WithTracer(tracer *go2sky.Tracer) applyOption {
	return func(o *options) {
		if tracer == nil {
			return
		}

		o.tracer = tracer
	}
}

// WithUrl allows to use customized tracer rather than the global one only.
func WithUrl(url string) applyOption {
	return func(o *options) {
		if url == "" {
			return
		}

		o.url = url
	}
}

func WithSqlParameters(logSqlParameters bool) applyOption {
	return func(o *options) {
		o.logSqlParameters = logSqlParameters
	}
}

func WithErrorTagHook(errorTagHook errorTagHook) applyOption {
	return func(o *options) {
		if errorTagHook == nil {
			return
		}

		o.errorTagHook = errorTagHook
	}
}

// errorTagHook will be called while gorm.DB got an error and we need a way to mark this error
// in current opentracing.Span. Of course, you can use sp.LogField in this hook, but it's not
// recommended to.
//
// mark an error tag in sp as default:
//
// sp.SetTag(sp.SetTag(_errorTagKey, true))
type errorTagHook func(sp *go2sky.Span, err error)

func defaultErrorTagHook(sp *go2sky.Span, err error) {
	(*sp).Tag(_errorTagKey, "true")
}

// tag called after operation
func tag(sp *go2sky.Span, db *gorm.DB, errorTagHook errorTagHook) {
	if err := db.Error; err != nil && nil != errorTagHook {
		errorTagHook(sp, err)
	}

	(*sp).Tag(go2sky.TagDBType, "mysql")
	(*sp).Tag(go2sky.TagDBInstance, db.Statement.DB.Name())
	(*sp).Tag(go2sky.TagDBStatement, db.Statement.Table)
}

// log called after operation
func sklog(sp *go2sky.Span, db *gorm.DB, verbose bool, logSqlVariables bool) {
	fields := make([]string, 0, 4)
	fields = appendSql(fields, db, logSqlVariables)
	fields = append(fields, _rowsAffectedLogKey+": "+strconv.FormatInt(db.Statement.RowsAffected, 10))

	// log error
	if err := db.Error; err != nil {
		fields = append(fields, "error: "+err.Error())
	}

	if verbose && db.Statement.Dest != nil {
		// DONE(@yeqown) fill result fields into span log
		// FIXED(@yeqown) db.Statement.Dest still be metatable now ?
		v, err := json.Marshal(db.Statement.Dest)
		if err == nil {
			fields = append(fields, _resultLogKey+": "+(*(*string)(unsafe.Pointer(&v))))
		} else {
			db.Logger.Error(db.Statement.Context, "could not marshal db.Statement.Dest: %v", err)
		}
	}

	(*sp).Log(time.Now(), fields...)
}

func appendSql(fields []string, db *gorm.DB, logSqlVariables bool) []string {
	if logSqlVariables {
		fields = append(fields, _sqlLogKey+": "+db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...))
	} else {
		fields = append(fields, _sqlLogKey+": "+db.Statement.SQL.String())
	}
	return fields
}

func (p skyWalkingPlugin) beforeCreate(db *gorm.DB) {
	p.injectBefore(db, _createOp)
}

func (p skyWalkingPlugin) after(db *gorm.DB) {
	p.extractAfter(db)
}

func (p skyWalkingPlugin) beforeUpdate(db *gorm.DB) {
	p.injectBefore(db, _updateOp)
}

func (p skyWalkingPlugin) beforeQuery(db *gorm.DB) {
	p.injectBefore(db, _queryOp)
}

func (p skyWalkingPlugin) beforeDelete(db *gorm.DB) {
	p.injectBefore(db, _deleteOp)
}

func (p skyWalkingPlugin) beforeRow(db *gorm.DB) {
	p.injectBefore(db, _rowOp)
}

func (p skyWalkingPlugin) beforeRaw(db *gorm.DB) {
	p.injectBefore(db, _rawOp)
}

// New constructs a new plugin based opentracing. It supports to trace all operations in gorm,
// so if you have already traced your servers, now this plugin will perfect your tracing job.
func NewSkyWalkingPlugin(opts ...applyOption) gorm.Plugin {
	dst := defaultOption()
	for _, apply := range opts {
		apply(dst)
	}

	return skyWalkingPlugin{
		//logResult: dst.logResult,
		opt: dst,
	}
}

func (p skyWalkingPlugin) Name() string {
	return "SkyWalking"
}

// Initialize registers all needed callbacks
func (p skyWalkingPlugin) Initialize(db *gorm.DB) (err error) {
	e := myError{
		errs: make([]string, 0, 12),
	}

	// create
	err = db.Callback().Create().Before("gorm:create").Register(_stageBeforeCreate.Name(), p.beforeCreate)
	e.add(_stageBeforeCreate, err)
	err = db.Callback().Create().After("gorm:create").Register(_stageAfterCreate.Name(), p.after)
	e.add(_stageAfterCreate, err)

	// update
	err = db.Callback().Update().Before("gorm:update").Register(_stageBeforeUpdate.Name(), p.beforeUpdate)
	e.add(_stageBeforeUpdate, err)
	err = db.Callback().Update().After("gorm:update").Register(_stageAfterUpdate.Name(), p.after)
	e.add(_stageAfterUpdate, err)

	// query
	err = db.Callback().Query().Before("gorm:query").Register(_stageBeforeQuery.Name(), p.beforeQuery)
	e.add(_stageBeforeQuery, err)
	err = db.Callback().Query().After("gorm:query").Register(_stageAfterQuery.Name(), p.after)
	e.add(_stageAfterQuery, err)

	// delete
	err = db.Callback().Delete().Before("gorm:delete").Register(_stageBeforeDelete.Name(), p.beforeDelete)
	e.add(_stageBeforeDelete, err)
	err = db.Callback().Delete().After("gorm:delete").Register(_stageAfterDelete.Name(), p.after)
	e.add(_stageAfterDelete, err)

	// row
	err = db.Callback().Row().Before("gorm:row").Register(_stageBeforeRow.Name(), p.beforeRow)
	e.add(_stageBeforeRow, err)
	err = db.Callback().Row().After("gorm:row").Register(_stageAfterRow.Name(), p.after)
	e.add(_stageAfterRow, err)

	// raw
	err = db.Callback().Raw().Before("gorm:raw").Register(_stageBeforeRaw.Name(), p.beforeRaw)
	e.add(_stageBeforeRaw, err)
	err = db.Callback().Raw().After("gorm:raw").Register(_stageAfterRaw.Name(), p.after)
	e.add(_stageAfterRaw, err)

	return e.toError()
}

type myError struct {
	errs []string
}

func (e *myError) add(stage operationStage, err error) {
	if err == nil {
		return
	}

	e.errs = append(e.errs, "stage="+stage.Name()+":"+err.Error())
}

func (e myError) toError() error {
	if len(e.errs) == 0 {
		return nil
	}

	return e
}

func (e myError) Error() string {
	return strings.Join(e.errs, ";")
}
