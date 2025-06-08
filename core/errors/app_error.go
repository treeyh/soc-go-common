package errors

import (
	"fmt"

	pkgErrors "github.com/pkg/errors"
)

var (
	_codes = make(map[int]ResultCode)
)

// ResultCode 调用返回状态
type ResultCode struct {
	code int

	message string

	args []interface{}

	error error
}

// https://github.com/pkg/errors/blob/master/errors.go

// AppError Application Error interface
type AppError interface {
	Error() string

	Code() int

	Message() string

	GetError() error

	GetMessage() string

	Args() []interface{}
}

func (rc *ResultCode) Error() string {
	return rc.error.Error()
}

// Code 返回状态编号
func (rc *ResultCode) Code() int {
	return rc.code
}

// Message 返回状态信息
func (rc *ResultCode) Message() string {
	if rc.args == nil || len(rc.args) == 0 {
		return rc.message
	}
	return fmt.Sprintf(rc.message, rc.args...)
}

// GetMessage 返回error
func (rc *ResultCode) GetMessage() string {
	return rc.message
}

// GetError 返回error
func (rc *ResultCode) GetError() error {
	return rc.error
}

// Args 返回error的Args
func (rc *ResultCode) Args() []interface{} {
	return rc.args
}

// SetCode 不开放  code readonly
// func (rc *ResultCode) SetCode(code int) {
// 	rc.code = code
// }

// SetMessage 不开放  message readonly
// func (rc *ResultCode) SetMessage(message string) {
// 	rc.message = message
// }

// NewResultCode 创建新的resultCode，code编号不允许与已有的重复，一般启动初始化时调用
func NewResultCode(code int, message string) ResultCode {
	if code < 0 {
		panic(fmt.Sprintf("result code: code %d must greater than zero", code))
	}
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("result code: %d already exist", code))
	}

	rci := ResultCode{
		code:    code,
		message: message,
	}
	_codes[code] = rci
	return rci
}

// NewResultCodeIgnoreDuplicate 创建新的resultCode 允许code重复，运行时调用
func NewResultCodeIgnoreDuplicate(code int, message string, args ...interface{}) ResultCode {
	rci := ResultCode{
		code:    code,
		message: message,
		args:    args,
	}
	return rci
}

// NewAppErrorByExistError 基于error创建应用错误，如果error为nil，则返回nil
func NewAppErrorByExistError(rc ResultCode, err error, args ...interface{}) AppError {
	if err == nil {
		return nil
	}
	rcc := &ResultCode{
		code:    rc.code,
		message: rc.message + "; err:" + err.Error(),
		error:   pkgErrors.Wrap(err, rc.message+"; err:"+err.Error()),
	}
	if len(args) > 0 {
		rcc.message = fmt.Sprintf(rc.message, args...) + "; err:" + err.Error()
		rcc.error = pkgErrors.Wrap(err, rcc.message)
		return rcc
	}
	if len(rc.args) > 0 {
		rcc.message = fmt.Sprintf(rc.message, rc.args...) + "; err:" + err.Error()
		rcc.error = pkgErrors.Wrap(err, rcc.message)
	}
	return rcc
}

func NewAppError(rc ResultCode, args ...interface{}) AppError {
	rcc := &ResultCode{
		code:    rc.code,
		message: rc.message,
		args:    args,
	}
	if len(args) > 0 {
		rcc.error = pkgErrors.New(fmt.Sprintf(rc.message, args...))
		return rcc
	}
	if len(rc.args) > 0 {
		rcc.args = rc.args
		rcc.error = pkgErrors.New(fmt.Sprintf(rc.message, rc.args...))
		return rcc
	}
	rcc.error = pkgErrors.New(rc.message)
	return rcc
}
