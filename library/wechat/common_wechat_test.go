package wechat

import (
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestWechatProxy_GetAccessToken(t *testing.T) {

	initWechatTestConfig()

	ctx := tests.GetNewContext()

	resp, err := GetProxy().getAccessToken(ctx)
	if err != nil {
		assert.NoError(t, err, "getAccessToken error. %+v.", err)
		return
	}

	respstr := json.ToJsonIgnoreError(resp)
	assert.Equal(t, resp.ExpiresIn, 7200 , "wechat getAccessToken error. %s", respstr)

	GetProxy().wechatConfig.AppId = GetProxy().wechatConfig.AppId + "_"

	resp, err = GetProxy().getAccessToken(ctx)
	log.Info(err.Message())
	log.Info(json.ToJsonIgnoreError(resp))

	assert.NoError(t, err , "getAccessToken error. %+v.", err)
}
