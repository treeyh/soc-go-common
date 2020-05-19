package wechat

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
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

	convey.Convey("log test", t, tests.TestStartUp(func() {

		resp, err := GetProxy().JsCode2Session(nil, "081GD5tq0egYKl1FfIrq0deGsq0GD5tm")
		if err != nil {
			convey.ShouldBeNil(err)
			return
		}

		convey.ShouldBeTrue(CheckErrCodeSucceed(resp.HttpStatus, resp.ErrCode))

		fmt.Println(json.ToJsonIgnoreError(resp))

		//result, err := GetProxy().DecryptEncryptedData(nil, resp.SessionKey,
		//	"HydtA2GK6zoPE7K4JN0HWU/uHI3sMkZ4F0dexYrg5pfBfANh42…Q1r9KxXUE8DqtIPEKRdvDShiZcbRF/vXxkRXswrjsWOmHFR0=",
		//	"SmblgkW2b5fDmZhXSxiukg==", true)

		result, err2 := GetProxy().DecryptEncryptedData(nil, resp.SessionKey,
			"aTossU7l+VWo7Eaz2+RsW3KecI0OSu4fZ/YyQwlb/SyOE8CK2ifbBXD7qb5GlG790OUbseKmZgF9e/AkXsMfaYY7D0bVjX9Bngh0z8CL+0eQFD7/Tjv6EoYLkk8LQogAiEUFlOMyv//8Xl8iDYVlxyYuYsSJQKNWrLdcJUhSruk14lig5pZdfD76rf12pdXmk66w5ERAP/7LIdTkOyEU3fJGVLhFteUvGdsVwwg1Q1Mczp6mzNSiDP2O6c+8X9lxpPVPk4rx+8gEYUxhZH97Wn1SxiBnfuHZMBPH5JVWZ5QGkO9SNa+AwBlnK5GM/m48SKx28YGDhTBcGqrQAMl0JdHhNqBOAROm09FURHV4eccnkFRMKPV6DbElDsW6cI5QIX9g/lsgnW3Vb4MhLC8RFCqQX6PVIQKDAWNsflIbDYa7ku5uUOQnqf4hYjtD54bJ10YM+c4kOwhPxllMkArnUwehPTknnAw+vtG2bQ37+GE=",
			"1S60BLqmpZuzljJG2/sPIA==", true)

		fmt.Println(err2)
		fmt.Println(result)
		fmt.Println(json.ToJsonIgnoreError(result))

	}, initWechatTestConfig))

}
