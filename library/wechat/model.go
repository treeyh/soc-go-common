package wechat

// WechatBaseResp 微信基础返回
type WechatBaseResp struct {
	ErrCode int64 `json:"errcode"`

	ErrMsg string `json:"errmsg"`

	HttpStatus int `json:"httpstatus"`
}

// WechatCode2SessionResp 微信code2Session返回
type WechatCode2SessionResp struct {
	WechatBaseResp

	OpenId string `json:"openid"`

	SessionKey string `json:"session_key"`

	UnionId string `json:"unionid"`
}
