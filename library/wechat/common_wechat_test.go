package wechat

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestWechatProxy_GetAccessToken(t *testing.T) {

	convey.Convey("TestWechatProxy_GetAccessToken test", t, tests.TestStartUp(func() {

		ctx := tests.GetNewContext()

		resp, err := GetProxy().getAccessToken(ctx)
		if err != nil {
			convey.So(err, convey.ShouldBeNil)
			return
		}

		log.Info(json.ToJsonIgnoreError(resp))
		//convey.So(CheckErrCodeSucceed(resp.HttpStatus, resp.ErrCode), convey.ShouldBeTrue)
		convey.So(resp.ExpiresIn, convey.ShouldEqual, 7200)

		GetProxy().wechatConfig.AppId = GetProxy().wechatConfig.AppId + "_"

		resp, err = GetProxy().getAccessToken(ctx)
		log.Info(err.Message())
		log.Info(json.ToJsonIgnoreError(resp))
		convey.So(err, convey.ShouldNotBeNil)

		convey.So(resp, convey.ShouldBeNil)

	}, initWechatTestConfig))

}
