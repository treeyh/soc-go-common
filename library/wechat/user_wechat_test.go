package wechat

import (
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestWechatProxy_GetOpenids(t *testing.T) {

	initWechatTestConfig()

	ctx := tests.GetNewContext()

	resp, err := GetProxy().GetOpenids(ctx, "")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Total>0)
	log.InfoCtx(ctx, json.ToJsonIgnoreError(resp))

	openid := resp.Data.Openid[0]

	userInfoReqs := make([]WechatUserInfoReq, 1)
	userInfoReqs[0] = WechatUserInfoReq{
		Openid: openid,
		Lang:   WechatLang_ZH_CN,
	}
	resp2, err := GetProxy().GetUserInfoBatch(ctx, userInfoReqs)
	assert.NoError(t, err)
	assert.NotNil(t, resp2)
	assert.True(t, len(resp2.UserInfoList) > 0)

	log.InfoCtx(ctx, json.ToJsonIgnoreError(resp2))

	_accessToken.AccessToken = _accessToken.AccessToken + "a"
	resp2, err = GetProxy().GetUserInfoBatch(ctx, userInfoReqs)
	assert.NoError(t, err)
	assert.NotNil(t, resp2)
	assert.True(t, len(resp2.UserInfoList) > 0)

	log.InfoCtx(ctx, json.ToJsonIgnoreError(resp2))
}
