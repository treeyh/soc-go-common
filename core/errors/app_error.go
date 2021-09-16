package errors

import (
	"fmt"
)

var (
	_codes = make(map[int]ResultCode)
)

// ResultCode 调用返回状态
type ResultCode struct {
	code int

	message string

	args []interface{}

	err error
}

// https://github.com/pkg/errors/blob/master/errors.go

// AppError Application Error interface
type AppError interface {
	Error() string

	Code() int

	Message() string

	GetError() error
}

func Wrap(err AppError, resultCode ResultCode) AppError {
	if err == nil {
		return NewAppError(resultCode)
	}

	return nil
}

func UnWrap() AppError {
	return nil
}

func (rc *ResultCode) Error() string {
	return rc.Message()
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

// GetError 返回error
func (rc *ResultCode) GetError() error {
	return rc.err
}

// SetCode 不开放  code readonly
// func (rc *ResultCode) SetCode(code int) {
// 	rc.code = code
// }

// SetMessage 不开放  message readonly
// func (rc *ResultCode) SetMessage(message string) {
// 	rc.message = message
// }

// NewResultCode 创建新的resultCode
// code编号不允许与已有的重复
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

// NewAppErrorByExistError 基于error创建应用错误，如果error为nil，则返回nil
func NewAppErrorByExistError(rc ResultCode, err error, e ...interface{}) AppError {
	if err == nil {
		return nil
	}

	rc.err = err
	rc.message += "; err:" + err.Error()
	if e != nil && len(e) == 0 {
		return &rc
	}

	rc.args = e
	return &rc
}

func NewAppError(rc ResultCode, e ...interface{}) AppError {
	if e == nil || (e != nil && len(e) == 0) {
		return &rc
	}
	rc.args = e
	return &rc
}
