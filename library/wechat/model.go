package wechat

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

type WechatUserInfoReq struct {
	Openid string     `json:"openid"`
	Lang   WechatLang `json:"lang"`
}

type WechatUserInfo struct {
	WechatErrorResp

	// Subscribe 用户是否订阅该公众号标识，值为 0 时，代表此用户没有关注该公众号，拉取不到其余信息。
	Subscribe int `json:"subscribe"`

	// Openid 用户的标识，对当前公众号唯一
	Openid string `json:"openid"`

	// Nickname 用户的昵称
	Nickname string `json:"nickname"`

	// Sex 用户的性别，值为 1 时是男性，值为 2 时是女性，值为 0 时是未知
	Sex int `json:"sex"`

	// Language 用户的语言，简体中文为 zh_CN
	Language string `json:"language"`

	// City 用户所在城市
	City string `json:"city"`

	// Province 用户所在省份
	Province string `json:"province"`

	// Country 用户所在国家
	Country string `json:"country"`

	// Headimgurl 用户头像，最后一个数值代表正方形头像大小（有 0、46、64、96、132 数值可选，0 代表 640*640 正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像 URL 将失效。
	Headimgurl string `json:"headimgurl"`

	// SubscribeTime 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	SubscribeTime int64 `json:"subscribe_time"`

	// Unionid 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Unionid string `json:"unionid"`

	// Remark 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Remark string `json:"remark"`

	// Groupid 用户所在的分组 ID（兼容旧的用户分组接口）
	Groupid int64 `json:"groupid"`

	// tagidList 用户被打上的标签 ID 列表
	tagidList []int `json:"tagid_list"`

	// SubscribeScene 返回用户关注的渠道来源，
	//     ADD_SCENE_SEARCH 公众号搜索，
	//     ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，
	//     ADD_SCENE_PROFILE_CARD 名片分享，
	//     ADD_SCENE_QR_CODE 扫描二维码，
	//     ADD_SCENE_PROFILE_LINK 图文页内名称点击，
	//     ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，
	//     ADD_SCENE_PAID 支付后关注，
	//     ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，
	//     ADD_SCENE_OTHERS 其他
	SubscribeScene string `json:"subscribe_scene"`

	// QrcScene 二维码扫码场景（开发者自定义）
	QrcScene int `json:"qr_scene"`

	// QrSceneStr 二维码扫码场景描述（开发者自定义）
	QrSceneStr string `json:"qr_scene_str"`
}

type WechatUserInfoList struct {
	WechatErrorResp

	UserInfoList []WechatUserInfo `json:"user_info_list"`
}

type WechatOpenidsResp struct {
	WechatErrorResp

	Total int64 `json:"total"`

	Count int `json:"count"`

	Data struct {
		Openid []string `json:"openid"`
	} `json:"data"`

	NextOpenid string `json:"next_openid"`
}

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
	MsgType WechatMsgType `json:"msgType" xml:"MsgType"`

	Event WechatEventType `json:"event" xml:"Event"`

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
