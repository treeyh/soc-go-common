package wechat

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestWechatProxy_SendTemplateMessage(t *testing.T) {

	convey.Convey("TestWechatProxy_SendTemplateMessage test", t, tests.TestStartUp(func() {

		templateMessageId := "th9QPtLjyAOpiu1-iDKzjUSb_hTKC5ZgF4h-_kznrpM"
		ctx := tests.GetNewContext()

		data := make(map[string]WechatTemplateMessageParamReq)
		data["first"] = WechatTemplateMessageParamReq{
			Value: "你好，Tree：",
			Color: "#173177",
		}
		data["keyword1"] = WechatTemplateMessageParamReq{
			Value: "keyword1：",
			Color: "#173177",
		}
		data["keyword2"] = WechatTemplateMessageParamReq{
			Value: "keyword222：",
			Color: "#173177",
		}
		data["remark"] = WechatTemplateMessageParamReq{
			Value: "remarkremark：",
			Color: "#173177",
		}

		req := &WechatTemplateMessageReq{
			ToUser:     "o8NgBv4OWPt5PVBJMnyU6vOj_i4s",
			TemplateId: templateMessageId,
			TopColor:   "#FF0000",
			Data:       data,
		}

		resp, err := GetProxy().SendTemplateMessage(ctx, req)
		convey.So(err, convey.ShouldBeNil)
		convey.So(resp.MsgId > 0, convey.ShouldBeTrue)

		log.InfoCtx(ctx, json.ToJsonIgnoreError(resp))

	}, initWechatTestConfig))
}
