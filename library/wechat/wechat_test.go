package wechat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/config"
	"testing"
)

func TestInitWeChatConfig(t *testing.T) {
	weChatConfigs := map[string]config.WeChatConfig{
		"default": config.WeChatConfig{
			AppId:               "AppId",
			AppSecret:           "AppSecret",
			Host:                "Host",
			Type:                "Type",
			Token:               "Token",
			EncodingAESKey:      "EncodingAESKey",
			MessageEncodingType: 1,
		},
		"test": config.WeChatConfig{
			AppId:               "1",
			AppSecret:           "2",
			Host:                "3",
			Type:                "4",
			Token:               "5",
			EncodingAESKey:      "6",
			MessageEncodingType: 7,
		},
	}

	InitWeChatConfig(weChatConfigs)

	fmt.Println(wechatProxys["default"].wechatConfig.AppId)
	fmt.Println(wechatProxys["test"].wechatConfig.AppId)
	assert.Equal(t, wechatProxys["default"].wechatConfig.Token, "Token")
	assert.Equal(t, wechatProxys["test"].wechatConfig.Token, "5")
}
