package aliyun

import (
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestAliYunProxy_SendSingleMail(t *testing.T) {

	initALiYunTestConfig()

	ctx := tests.GetNewContext()

	accountName := "no-reply@service.mail.yuekai.top"
	replyToAddress := false
	toAddress := []string{"cr@ejyi.com", "tree@ejyi.com"}
	fromAlias := "FromAlias"
	subject := "subject"
	htmlBody := "htmlBody-4"

	resp, err := GetProxy().SendSingleMail(ctx, accountName, toAddress, replyToAddress, fromAlias, subject, htmlBody)

	assert.NoError(t, err)
	assert.True(t, resp.IsSuccess())

	log.InfoCtx(ctx, resp.IsSuccess())
	log.InfoCtx(ctx, resp.GetHttpContentString())
	log.InfoCtx(ctx, resp.GetHttpStatus())
	log.InfoCtx(ctx, resp.GetHttpHeaders())
}
