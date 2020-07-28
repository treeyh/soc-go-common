package wechat

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestWechatProxy_GetOpenids(t *testing.T) {

	convey.Convey("TestWechatProxy_GetOpenids test", t, tests.TestStartUp(func() {

		ctx := tests.GetNewContext()

		resp, err := GetProxy().GetOpenids(ctx, "")

		convey.So(err, convey.ShouldBeNil)
		convey.So(resp, convey.ShouldNotBeNil)
		convey.So(resp.Total > 0, convey.ShouldBeTrue)
		log.InfoCtx(ctx, json.ToJsonIgnoreError(resp))

		openid := resp.Data.Openid[0]

		userInfoReqs := make([]WechatUserInfoReq, 1)
		userInfoReqs[0] = WechatUserInfoReq{
			Openid: openid,
			Lang:   WechatLang_ZH_CN,
		}
		resp2, err := GetProxy().GetUserInfoBatch(ctx, userInfoReqs)
		convey.So(err, convey.ShouldBeNil)
		convey.So(resp2, convey.ShouldNotBeNil)
		convey.So(len(resp2.UserInfoList) > 0, convey.ShouldBeTrue)
		log.InfoCtx(ctx, json.ToJsonIgnoreError(resp2))

	}, initWechatTestConfig))
}
