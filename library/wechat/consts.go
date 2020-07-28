package wechat

import "reflect"

const accessTokenKey = "access_token"

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

type (
	WechatMsgType   string
	WechatEventType string
	WechatLang      string
)

// 消息类型
const (
	MSGTYPE_TEXT       WechatMsgType = "text"
	MSGTYPE_IMAGE      WechatMsgType = "image"
	MSGTYPE_VOICE      WechatMsgType = "voice"
	MSGTYPE_VIDEO      WechatMsgType = "video"
	MSGTYPE_SHORTVIDEO WechatMsgType = "shortvideo"
	MSGTYPE_LOCATION   WechatMsgType = "location"
	MSGTYPE_LINK       WechatMsgType = "link"
	MSGTYPE_EVENT      WechatMsgType = "event"
)

// 事件推送类型
const (
	EVENT_SUBSCRIBE   WechatEventType = "subscribe"
	EVENT_UNSUBSCRIBE WechatEventType = "unsubscribe"
	EVENT_SCAN        WechatEventType = "scan"
	EVENT_LOCATION    WechatEventType = "location"
	EVENT_CLICK       WechatEventType = "click"
)

// 语言
const (
	WechatLang_ZH_CN WechatLang = "zh_CN"

	WechatLang_ZH_TW WechatLang = "zh_TW"

	WechatLang_EN WechatLang = "en"
)

type MenuButtonType string

const (
	// MenuButtonTypeClick 表示点击类型
	MenuButtonTypeClick MenuButtonType = "click"

	// MenuButtonTypeView 表示网页类型
	MenuButtonTypeView MenuButtonType = "view"

	// MenuButtonTypeMiniProgram 表示小程序类型
	MenuButtonTypeMiniProgram MenuButtonType = "miniprogram"
)

const (
	// SubscribeSceneType_ADD_SCENE_SEARCH 公众号搜索
	SubscribeSceneType_ADD_SCENE_SEARCH = "ADD_SCENE_SEARCH"

	// SubscribeSceneType_ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移
	SubscribeSceneType_ADD_SCENE_ACCOUNT_MIGRATION = "ADD_SCENE_ACCOUNT_MIGRATION"

	// SubscribeSceneType_ADD_SCENE_PROFILE_CARD 名片分享
	SubscribeSceneType_ADD_SCENE_PROFILE_CARD = "ADD_SCENE_PROFILE_CARD"

	// SubscribeSceneType_ADD_SCENE_QR_CODE 扫描二维码
	SubscribeSceneType_ADD_SCENE_QR_CODE = "ADD_SCENE_QR_CODE"

	// SubscribeSceneType_ADD_SCENE_PROFILE_LINK 图文页内名称点击
	SubscribeSceneType_ADD_SCENE_PROFILE_LINK = "ADD_SCENE_PROFILE_LINK"

	// SubscribeSceneType_ADD_SCENE_PROFILE_ITEM 图文页右上角菜单
	SubscribeSceneType_ADD_SCENE_PROFILE_ITEM = "ADD_SCENE_PROFILE_ITEM"

	// SubscribeSceneType_ADD_SCENE_PAID 支付后关注
	SubscribeSceneType_ADD_SCENE_PAID = "ADD_SCENE_PAID"

	// SubscribeSceneType_ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告
	SubscribeSceneType_ADD_SCENE_WECHAT_ADVERTISEMENT = "ADD_SCENE_WECHAT_ADVERTISEMENT"

	// SubscribeSceneType_ADD_SCENE_OTHERS 其他
	SubscribeSceneType_ADD_SCENE_OTHERS = "ADD_SCENE_OTHERS"
)
