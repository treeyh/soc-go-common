package wechat

// 消息类型
const (
	MSGTYPE_TEXT       = "text"
	MSGTYPE_IMAGE      = "image"
	MSGTYPE_VOICE      = "voice"
	MSGTYPE_VIDEO      = "video"
	MSGTYPE_SHORTVIDEO = "shortvideo"
	MSGTYPE_LOCATION   = "location"
	MSGTYPE_LINK       = "link"
	MSGTYPE_EVENT      = "event"
)

// 事件推送类型
const (
	EVENT_SUBSCRIBE   = "subscribe"
	EVENT_UNSUBSCRIBE = "unsubscribe"
	EVENT_SCAN        = "scan"
	EVENT_LOCATION    = "location"
	EVENT_CLICK       = "click"
)

// 基础数据
type BaseData struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
}

type CDATA struct {
	Content string `xml:",cdata"`
}

// 文本消息
type TextMsg struct {
	BaseData
	Content string `xml:"Content"`
	MsgId   string `xml:"MsgId"`
}

type ReplyTextMsg struct {
	XMLName      string `xml:"xml"`
	ToUserName   CDATA  `xml:"ToUserName"`
	FromUserName CDATA  `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      CDATA  `xml:"MsgType"`
	Content      CDATA  `xml:"Content"`
}

// 关注 / 取关事件
type SubscribeEvent struct {
	BaseData
	Event string `xml:"Event"`
}

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

type WechatAccessTokenResp struct {
	WechatBaseResp

	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
