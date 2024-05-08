package json

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	jsonextra "github.com/json-iterator/go/extra"
	"github.com/treeyh/soc-go-common/core/consts"
)

type TimeDecoder struct{}
type TimeEncoder struct {
	precision time.Duration
}

var loc, _ = time.LoadLocation("Local")
var log = logger.Logger()

func (td *TimeDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	str := iter.ReadString()

	mayBlank, _ := time.Parse(consts.AppTimeFormat, str)
	now, err := time.ParseInLocation(consts.AppTimeFormat, str, loc)

	if err != nil {
		mayBlank2, err2 := time.Parse(consts.AppSystemTimeFormat8, str)
		if err2 == nil {
			*((*time.Time)(ptr)) = mayBlank2
		} else {
			*((*time.Time)(ptr)) = time.Unix(0, 0)
		}
	} else if mayBlank.IsZero() {
		*((*time.Time)(ptr)) = mayBlank
	} else {
		*((*time.Time)(ptr)) = now
	}
}

func (codec *TimeEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	ts := *((*time.Time)(ptr))
	return ts.IsZero()
}

func (codec *TimeEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := *((*time.Time)(ptr))
	if !ts.IsZero() {
		//timestamp := ts.Unix()
		//tm := time.Unix(timestamp, 0)
		//format := tm.Format(consts.AppTimeFormat)

		format := ts.Format(consts.AppSystemTimeFormat8)
		stream.WriteString(format)
	} else {
		mayBlank, _ := time.Parse(consts.AppSystemTimeFormat8, consts.BlankString)
		stream.WriteString(mayBlank.Format(consts.AppSystemTimeFormat8))
	}
}

var json = jsoniter.ConfigFastest

func init() {
	jsoniter.RegisterTypeEncoder("time.Time", &TimeEncoder{})
	jsoniter.RegisterTypeDecoder("time.Time", &TimeDecoder{})

	// 容忍字符串和数字互转
	jsonextra.RegisterFuzzyDecoders()
}

func ToJson(obj interface{}) (string, errors.AppError) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return "", errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return string(bs), nil
}

func ToJsonCtx(ctx context.Context, obj interface{}) (string, errors.AppError) {
	bs, err := json.Marshal(obj)
	if err != nil {
		log.ErrorCtx2(ctx, err, "json encode fail.")
		return "", errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return string(bs), nil
}

func ToJsonIgnoreError(obj interface{}) string {
	jstr, _ := ToJson(obj)
	return jstr
}

func FromJson(jsonStr string, obj interface{}) errors.AppError {
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return nil
}

func FromJsonCtx(ctx context.Context, jsonStr string, obj interface{}) errors.AppError {
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		log.ErrorCtx2(ctx, err, "json decode fail.")
		return errors.NewAppErrorByExistError(errors.RedisOperationFail, err)
	}
	return nil
}
