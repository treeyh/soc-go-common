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

// WechatCode2AccessTokenResp 微信code获取accessToken返回
type WechatCode2AccessTokenResp struct {
	WechatErrorResp

	// AccessToken 接口调用凭证
	AccessToken string `json:"access_token"`

	// ExpiresIn access_token 接口调用凭证超时时间，单位（秒）
	ExpiresIn int `json:"expires_in"`

	// RefreshToken 用户刷新 access_token
	RefreshToken string `json:"refresh_token"`

	// Openid 授权用户唯一标识
	Openid string `json:"openid"`

	// Scope 用户授权的作用域，使用逗号（,）分隔
	Scope string `json:"scope"`

	// Unionid 当且仅当该移动应用已获得该用户的 userinfo 授权时，才会出现该字段
	Unionid string `json:"unionid"`
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

	// Privilege 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Privilege []string `json:"privilege"`

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

type WechatTemplateMessageParamReq struct {
	Value string `json:"value,omitempty"`

	Color string `json:"color,omitempty"`
}

type WechatTemplateMessageReq struct {
	ToUser string `json:"touser,omitempty"`

	TemplateId string `json:"template_id,omitempty"`

	Url string `json:"url,omitempty"`

	TopColor string `json:"topcolor,omitempty"`

	Data map[string]WechatTemplateMessageParamReq `json:"data,omitempty"`
}

type WechatMessageResp struct {
	WechatErrorResp

	MsgId int64 `json:"msgid"`
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
	MsgType WechatMsgType `json:"msgType,omitempty" xml:"MsgType,omitempty"`

	Event WechatEventType `json:"event,omitempty" xml:"Event,omitempty"`

	EventKey string `json:"eventKey,omitempty" xml:"EventKey,omitempty"`

	MenuID string `json:"menuID,omitempty" xml:"MenuID,omitempty"`

	MsgId int64 `xml:"MsgId"        json:"msgId"` // request

	Content string `xml:"Content,omitempty"      json:"content,omitempty"` // request

	MediaId string `xml:"MediaId,omitempty"      json:"mediaId,omitempty"` // request

	PicURL string `xml:"PicUrl,omitempty"       json:"picUrl,omitempty"` // request

	Format string `xml:"Format,omitempty"       json:"format,omitempty"` // request

	Recognition string `xml:"Recognition,omitempty"  json:"recognition,omitempty"` // request

	ThumbMediaId string `xml:"ThumbMediaId,omitempty" json:"thumbMediaId,omitempty"` // request

	LocationX float64 `xml:"Location_X,omitempty"   json:"locationX,omitempty"` // request

	LocationY float64 `xml:"Location_Y,omitempty"   json:"locationY,omitempty"` // request

	Scale int `xml:"Scale,omitempty"        json:"scale,omitempty"` // request

	Label string `xml:"Label,omitempty"        json:"label,omitempty"` // request

	Title string `xml:"Title,omitempty"        json:"title,omitempty"` // request

	Description string `xml:"Description,omitempty"  json:"description,omitempty"` // request

	URL string `xml:"Url,omitempty"          json:"url,omitempty"` // request

	Ticket string `xml:"Ticket,omitempty"       json:"ticket,omitempty"` // request

	Latitude float64 `xml:"Latitude,omitempty"     json:"latitude,omitempty"` // request

	Longitude float64 `xml:"Longitude,omitempty"    json:"longitude,omitempty"` // request

	Precision float64 `xml:"Precision,omitempty"    json:"precision,omitempty"` // request

	Status string `json:"status,omitempty" xml:"Status,omitempty"`
}

type WechatArticle struct {
	Title       string `xml:"Title,omitempty"       json:"title,omitempty"`       // 图文消息标题
	Description string `xml:"Description,omitempty" json:"description,omitempty"` // 图文消息描述
	PicURL      string `xml:"PicUrl,omitempty"      json:"picUrl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         string `xml:"Url,omitempty"         json:"url,omitempty"`         // 点击图文消息跳转链接
}

type WechatResponseBaseMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// ToUserName 开发者 微信号
	ToUserName string `json:"toUserName" xml:"ToUserName"`

	// FromUserName 发送方帐号（一个 OpenID）
	FromUserName string `json:"fromUserName" xml:"FromUserName"`

	// CreateTime 消息创建时间 （整型）
	CreateTime int64 `json:"createTime" xml:"CreateTime"`

	// MsgType 消息类型，event
	MsgType WechatMsgType `json:"msgType" xml:"MsgType"`
}

// WechatResponseTextMsg 微信回复文本消息
type WechatResponseTextMsg struct {
	WechatResponseBaseMsg

	// Content 回复的消息内容（换行：在 content 中能够换行，微信客户端就支持换行显示）
	Content string `json:"content,omitempty" xml:"Content,omitempty"`
}

// WechatResponseImageMsg 微信回复图片消息
type WechatResponseImageMsg struct {
	WechatResponseBaseMsg

	// Image 图片消息
	Image struct {
		MediaId string `json:"mediaId,omitempty" xml:"MediaId,omitempty"`
	} `json:"image,omitempty" xml:"Image,omitempty"`
}

// WechatResponseVoiceMsg 微信回复音频消息
type WechatResponseVoiceMsg struct {
	WechatResponseBaseMsg

	Voice struct {
		MediaId string `json:"mediaId,omitempty" xml:"MediaId,omitempty"`
	} `json:"voice,omitempty" xml:"Voice,omitempty"`
}

// WechatResponseVideoMsg 微信回复视频消息
type WechatResponseVideoMsg struct {
	WechatResponseBaseMsg

	Video *struct {
		MediaId     string `json:"mediaId,omitempty" xml:"MediaId,omitempty"`
		Title       string `json:"title,omitempty" xml:"Title,omitempty"`
		Description string `json:"description,omitempty" xml:"Description,omitempty"`
	} `json:"video,omitempty" xml:"Video,omitempty"`
}

// WechatResponseMusicMsg 微信回复图片消息
type WechatResponseMusicMsg struct {
	WechatResponseBaseMsg

	Music *struct {
		MusicUrl    string `json:"musicUrl,omitempty" xml:"MusicUrl,omitempty"`
		HQMusicUrl  string `json:"hQMusicUrl,omitempty" xml:"HQMusicUrl,omitempty"`
		Title       string `json:"title,omitempty" xml:"Title,omitempty"`
		Description string `json:"description,omitempty" xml:"Description,omitempty"`

		ThumbMediaId string `json:"thumbMediaId,omitempty" xml:"ThumbMediaId,omitempty"`
	} `json:"music,omitempty" xml:"Music,omitempty"`
}

// WechatResponseArticlesMsg 微信回复图片图文消息
type WechatResponseArticlesMsg struct {
	WechatResponseBaseMsg

	Articles []WechatArticle `xml:"Articles>item,omitempty" json:"Articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}
