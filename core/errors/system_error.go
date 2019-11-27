package errors

import (
	"fmt"
	"strconv"
)

var (
	_codes = make(map[int]ResultCode)
)

// ResultCode 调用返回状态
type ResultCode struct {
	code int

	message string

	args []interface{}
}

// System Error interface
type SystemError interface {
	Error() string

	Code() int

	Message() string
}

func (rc *ResultCode) Error() string {
	return "{\"code\":" + strconv.Itoa(rc.code) + ",\"message\":\"" + rc.message + "\"}"
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
	return fmt.Sprintf(rc.message, (rc.args)...)
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

func NewSystemError(rc ResultCode, e ...interface{}) SystemError {
	if e == nil || (e != nil && len(e) == 0) {
		return &rc
	}

	rc.args = e
	return &rc
}

