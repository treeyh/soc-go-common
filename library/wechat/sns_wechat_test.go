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
	InitWeChatConfig(&wecmap)
}

func TestDecrypt(t *testing.T) {

	convey.Convey("log test", t, tests.TestStartUp(func() {

		resp, err := GetProxy().JsCode2Session(nil, "011cwx6d2xeheI0Y8Cad2twK6d2cwx66")
		if err != nil {
			convey.ShouldBeNil(err)
			return
		}

		convey.ShouldBeTrue(CheckErrCodeSucceed(resp.HttpStatus, resp.ErrCode))

		fmt.Println(json.ToJsonIgnoreError(resp))

		result, err := GetProxy().DecryptEncryptedData(nil, resp.SessionKey,
			"Q8PRihnWUuAYSqkLb/PB3eBb/23uiAM1wwhd826pHek8vNLMtXJ9WBLPS4tTRucb5WLX+RGe1zWT49grPoRtGaq/Goiqv+TgLxhxdZwm63R/N5ShfAwrmqKgkoAMknPxV4J7D+wJnn6Fh/R0se8NFIlgyrywmeXx5vPjhLeQB7gMh1SneD6fsOLOISMMHFsPKWG9WXeeW4PgzjPJGNjVvB3CxZCeMpZvk1jyS7ZfVsLEQrmdlthRZAgFWOmIn0KaEk9n4NErW49k9wrId8EELIptC0NoteVE+39FZFowMrsuWkEYgZk/9QyLZn8Et2FHmBEYYmuYx1fbung5WkItoYPMDV+ofKewfdO2jW8/6PRYG+6XsmgHEuyxJP0VzK2FIdru+Cge8NtI7E1C1FN5YbwtSRcFLion/mdAb7SWUpYusRiEgc0NNgqdK1h4hValseLG7Kr7bGrC3YzN5+qEvV2Jy86VxiNWg+hGLHEv7dE=",
			"UaR85svk3mSqZI5VhKIrcw==", false)

		fmt.Println(json.ToJsonIgnoreError(result))

	}, initWechatTestConfig))

}
