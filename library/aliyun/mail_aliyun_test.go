package aliyun

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestAliYunProxy_SendSingleMail(t *testing.T) {

	convey.Convey("TestAliYunProxy_SendSingleMail test", t, tests.TestStartUp(func() {

		ctx := tests.GetNewContext()

		accountName := "no-reply@service.mail.yuekai.top"
		replyToAddress := false
		toAddress := []string{"cr@ejyi.com", "tree@ejyi.com"}
		fromAlias := "FromAlias"
		subject := "subject"
		htmlBody := "htmlBody-4"

		resp, err := GetProxy().SendSingleMail(ctx, accountName, toAddress, replyToAddress, fromAlias, subject, htmlBody)
		convey.So(err, convey.ShouldBeNil)
		convey.So(resp.IsSuccess(), convey.ShouldBeTrue)

		log.InfoCtx(ctx, resp.IsSuccess())
		log.InfoCtx(ctx, resp.GetHttpContentString())
		log.InfoCtx(ctx, resp.GetHttpStatus())
		log.InfoCtx(ctx, resp.GetHttpHeaders())

	}, initALiYunTestConfig))
}
