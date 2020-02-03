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

		resp, err := GetProxy().JsCode2Session(nil, "081wqyl81ySIgP1vbcl81qXpl81wqylV")
		if err != nil {
			convey.ShouldBeNil(err)
			return
		}

		convey.ShouldBeTrue(CheckErrCodeSucceed(resp.HttpStatus, resp.ErrCode))

		fmt.Println(json.ToJsonIgnoreError(resp))

		result, err := GetProxy().DecryptEncryptedData(nil, resp.SessionKey,
			"yWPWP6HKLdkunn9X9CtYL1TTOhQV5u+KyRY6Bc6dISmP7YwSDzb5A7LgLIMEpprC5eD8CzAJapcBApg7gGQYYM6NgCKh6h14NedzPWLhl1k1CYrtm6h+Eu1Icc4J0f8Xbpp5H6RSp2k+H1m+qbOorknGyNWQOC/o9FmjKML5QjMOGZ5FLcB33wtzMMyM2I1BhzCl/4lVwfIvgaPOKTvfETvXny59oAGuFMfMNgQxCj8dR9LctnvQlM+Evv8zRhgDwx9Ls4Ltn1I1Du+kyh4szHBpuCFUGipZdc93NH0dsfHaExY7DCGoTevEg+dEE3iJVmulo+eBN2qVnrdV5xs4qAdgLyplFjkvFAWllq55MvLUOAzV3TyiLFggGqJWRI0AcXoCwgjwWWH2jsvnlFLO+4dUfUdlo0GK9a8SzFTa7v8DB/cdH9lXVhNGrtoxsQuenkJIGmms1OpjT5C4t26ixkmSRVJVSWD0TfCF7XURhn4=",
			"TUC+krAFPmj+Pf5OuJjn7A==", true)

		fmt.Println(result)
		fmt.Println(json.ToJsonIgnoreError(result))

	}, initWechatTestConfig))

}
