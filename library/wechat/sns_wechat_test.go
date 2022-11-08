package wechat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/utils/file"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"github.com/treeyh/soc-go-common/tests"
	"path"
	"testing"
)

func initWechatTestConfig() {

	wechatJson, _ := file.ReadSmallFile(path.Join(file.GetCurrentPath(), "..", "..", "tests", "wechat.log"))

	wechatConfig := &config.WeChatConfig{}
	err := json.FromJson(*wechatJson, wechatConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	wecmap := make(map[string]config.WeChatConfig)
	wecmap[_defaultConfigName] = *wechatConfig
	InitWeChatConfig(wecmap)
}

func TestDecrypt(t *testing.T) {

	initWechatTestConfig()

	resp, err := GetProxy().JsCode2Session(nil, "0519YF000PHPRO1QkK100CAynK29YF0V")
	if err != nil {
		assert.NoError(t, err)
		return
	}

	assert.True(t, CheckErrCodeSucceed(200, resp.ErrCode))

	// {\"errcode\":0,\"errmsg\":\"\",\"openid\":\"oO7wo5EvObtcivnQtHxs8kmbn30A\",\"session_key\":\"WdU2sNq8VFlvXgAgFcUkSg==\",\"unionid\":\"ot07f0ZtRZ2A_sxv_jtYhpcbhqec\"}
	log.Info(json.ToJsonIgnoreError(resp))

	//result, err := GetProxy().DecryptEncryptedData(nil, resp.SessionKey,
	//	"HydtA2GK6zoPE7K4JN0HWU/uHI3sMkZ4F0dexYrg5pfBfANh42…Q1r9KxXUE8DqtIPEKRdvDShiZcbRF/vXxkRXswrjsWOmHFR0=",
	//	"SmblgkW2b5fDmZhXSxiukg==", true)

	result, err2 := GetProxy().DecryptEncryptedData(nil, resp.SessionKey,
		"IRNdYNi84ZEk/yZwBLsnrHcH4M/2FD5KMXphZZpza9zlokR87hnQHoCGksovkZC+KGPAe9HRtSTg2ifWJ31jHxnsSn6XJYmWiZ7+de2MINVIycVDgqTG6niJDp4OYPiM1lpsr5JZUjcDmH1tTC62rZLlknGwzeCZZ9wD1Jp32W1bovLFvlbjNwYMWpk5OBSP2vnzC9yqPgvGLuxhYTEhxiz6BiZPAPaj7+7XBzv+gjrmiS7yrnAZIUkSHnvwVA0zzSW1lEr/NXkZx60g9JYFSP38ViIDXLvv8UzEHBZ5FXNFsbsI1fKk7ePdk9HMom/Ewv+7BV44MjR+K3EEPQMDsYu/b9WGZ0w6pHZh8upRE6kEQBVhKOhmYhZaClxgG7rTLZ+cSyIpaFQEar/SQavinh4+4nCzprJn0fM4YGwkhE4=",
		"dLwV9296bU7nPGkzldrrjQ==", true)

	log.Info(err2)
	log.Info(result)
	log.Info(json.ToJsonIgnoreError(result))

}

func TestAppLogin(t *testing.T) {

	ctx := tests.GetNewContext()

	initWechatTestConfig()

	resp, err := GetProxy().GetAccessTokenByCode(ctx, "0115N6000zeoSO1afT3000iPhm45N60o")
	if err != nil {
		assert.NoError(t, err)
		return
	}

	assert.True(t, CheckErrCodeSucceed(200, resp.ErrCode))

	log.Info(json.ToJsonIgnoreError(resp))

	result, err2 := GetProxy().GetSnsUserInfo(nil, resp.AccessToken, resp.Openid, "zh_CN")

	log.Info(err2)
	log.Info(result)
	log.Info(json.ToJsonIgnoreError(result))

}
