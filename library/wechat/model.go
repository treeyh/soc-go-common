package wechat

type (
	MsgType   string
	EventType string
)

// 消息类型
const (
	MSGTYPE_TEXT       MsgType = "text"
	MSGTYPE_IMAGE      MsgType = "image"
	MSGTYPE_VOICE      MsgType = "voice"
	MSGTYPE_VIDEO      MsgType = "video"
	MSGTYPE_SHORTVIDEO MsgType = "shortvideo"
	MSGTYPE_LOCATION   MsgType = "location"
	MSGTYPE_LINK       MsgType = "link"
	MSGTYPE_EVENT      MsgType = "event"
)

// 事件推送类型
const (
	EVENT_SUBSCRIBE   EventType = "subscribe"
	EVENT_UNSUBSCRIBE EventType = "unsubscribe"
	EVENT_SCAN        EventType = "scan"
	EVENT_LOCATION    EventType = "location"
	EVENT_CLICK       EventType = "click"
)

// WechatErrorResp 微信基础返回
type WechatErrorResp struct {
	ErrCode int64 `json:"errcode"`

	ErrMsg string `json:"errmsg"`
}

// WechatCode2SessionResp 微信code2Session返回
type WechatCode2SessionResp struct {
	WechatErrorResp

	OpenId string `json:"openid"`

	SessionKey string `json:"session_key"`

	UnionId string `json:"unionid"`
}

type WechatAccessTokenResp struct {
	WechatErrorResp

	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type MenuButtonType string

const (
	// MenuButtonTypeClick 表示点击类型
	MenuButtonTypeClick MenuButtonType = "click"

	// MenuButtonTypeView 表示网页类型
	MenuButtonTypeView MenuButtonType = "view"

	// MenuButtonTypeMiniProgram 表示小程序类型
	MenuButtonTypeMiniProgram MenuButtonType = "miniprogram"
)

// WechatMenu 微信公众号菜单对象
type WechatMenu struct {
	Button []WechatMenuButton `json:"button"`
}

// WechatMenuButton 微信公众号菜单按钮对象
type WechatMenuButton struct {
	Type MenuButtonType `json:"type"`
	Name string         `json:"name"`
	Key  string         `json:"key"`

	Url       string             `json:"url"`
	MediaId   string             `json:"media_id"`
	Appid     string             `json:"appid"`
	Pagepath  string             `json:"pagepath"`
	SubButton []WechatMenuButton `json:"sub_button"`
}

// WechatReqestMsg 微信请求消息
type WechatReqestMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// ToUserName 开发者 微信号
	ToUserName string `json:"toUserName" xml:"ToUserName"`

	// FromUserName 发送方帐号（一个 OpenID）
	FromUserName string `json:"fromUserName" xml:"FromUserName"`

	// CreateTime 消息创建时间 （整型）
	CreateTime int64 `json:"createTime" xml:"CreateTime"`

	// MsgType 消息类型，event
	MsgType MsgType `json:"msgType" xml:"MsgType"`

	Event EventType `json:"event" xml:"Event"`

	EventKey string `json:"eventKey" xml:"EventKey"`

	MenuID string `json:"menuID" xml:"MenuID"`

	MsgId int64 `xml:"MsgId"        json:"msgId"` // request

	Content string `xml:"Content"      json:"content"` // request

	MediaId string `xml:"MediaId"      json:"mediaId"` // request

	PicURL string `xml:"PicUrl"       json:"picUrl"` // request

	Format string `xml:"Format"       json:"format"` // request

	Recognition string `xml:"Recognition"  json:"recognition"` // request

	ThumbMediaId string `xml:"ThumbMediaId" json:"thumbMediaId"` // request

	LocationX float64 `xml:"Location_X"   json:"locationX"` // request

	LocationY float64 `xml:"Location_Y"   json:"locationY"` // request

	Scale int `xml:"Scale"        json:"scale"` // request

	Label string `xml:"Label"        json:"label"` // request

	Title string `xml:"Title"        json:"title"` // request

	Description string `xml:"Description"  json:"description"` // request

	URL string `xml:"Url"          json:"url"` // request

	Ticket string `xml:"Ticket"       json:"ticket"` // request

	Latitude float64 `xml:"Latitude"     json:"latitude"` // request

	Longitude float64 `xml:"Longitude"    json:"longitude"` // request

	Precision float64 `xml:"Precision"    json:"precision"` // request

}
