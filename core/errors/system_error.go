package errors

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

var (
	_codes = make(map[int]ResultCode)
)

// result code struct
type ResultCode struct {
	code int

	message string

	langCode string
}

// System Error interface
type SystemError interface {
	Error() string

	Code() int

	Message() string

	LangCode() string

	Equal(error) bool
}

func (sei ResultCode) Error() string {
	return "{\"code\":" + strconv.Itoa(sei.code) + ",\"message\":\"" + sei.message + "\"}"
}

func (sei ResultCode) Code() int {
	return sei.code
}

func (sei ResultCode) Message() string {
	return sei.message
}

func (sei ResultCode) LangCode() string {
	return sei.langCode
}

func (sei *ResultCode) SetCode(code int) {
	sei.code = code
}

func (sei *ResultCode) SetMessage(message string) {
	sei.message = message
}

func (sei *ResultCode) SetLangCode(langCode string) {
	sei.langCode = langCode
}

// Equal check error is ResultCode
func (sei ResultCode) Equal(err error) bool {
	return EqualError(sei, err)
}

// GetResultCodeByCode get ResultCode by code
func GetResultCodeByCode(code int) ResultCode {

	if v, ok := _codes[code]; ok {
		return v
	} else {
		return SystemErr
	}
}

func getResultCodeByString(e string) ResultCode {
	if e == "" {
		return OK
	}

	i, err := strconv.Atoi(e)
	if err != nil {
		return SystemErr
	}
	if v, ok := _codes[i]; ok {
		return v
	} else {
		return SystemErr
	}
}

func convertSystemError(e error) SystemError {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(SystemError)
	if ok {
		return ec
	}
	return getResultCodeByString(e.Error())
}

func EqualError(code ResultCode, err error) bool {
	return convertSystemError(err).Code() == code.Code()
}

//add system ResultCode
func AddResultCode(code int, message string, langCode string) ResultCode {
	if code < 0 {
		panic(fmt.Sprintf("result code: code %d must greater than zero", code))
	}
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("result code: %d already exist", code))
	}

	rci := ResultCode{
		code:     code,
		message:  message,
		langCode: langCode,
	}

	_codes[code] = rci
	return rci
}

func BuildSystemError(ResultCode ResultCode, e ...error) SystemError {
	if e == nil || (e != nil && len(e) == 0) {
		return ResultCode
	}
	if e != nil {
		if reflect.TypeOf(e[0]).Name() == "ResultCode" {
			v := e[0].(SystemError)
			return v
		} else {
			ResultCode.message += "," + e[0].Error()
		}
	}
	return ResultCode
}

func BuildSystemErrorWithMessage(ResultCode ResultCode, message string) SystemError {
	ResultCode.message += message
	return ResultCode
}
