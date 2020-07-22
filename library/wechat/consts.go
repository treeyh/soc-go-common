package wechat

import "reflect"

const (
	ErrCodeOK                 = 0
	ErrCodeInvalidCredential  = 40001 // access_token 过期错误码
	ErrCodeAccessTokenExpired = 42001 // access_token 过期错误码(maybe!!!)
)

var (
	errorType      = reflect.TypeOf(WechatErrorResp{})
	errorZeroValue = reflect.Zero(errorType)
)

const (
	errorErrCodeIndex = 0
	errorErrMsgIndex  = 1
)
