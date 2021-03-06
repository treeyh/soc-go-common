package aliyun

import (
	"fmt"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/utils/file"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"path"
)

func initALiYunTestConfig() {

	aLiYunJson, _ := file.ReadSmallFile(path.Join(file.GetCurrentPath(), "..", "..", "tests", "aliyun.log"))

	wechatConfig := &config.ALiYunConfig{}
	err := json.FromJson(*aLiYunJson, wechatConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	alymap := make(map[string]config.ALiYunConfig)
	alymap[_defaultConfigName] = *wechatConfig
	InitALiYunConfig(alymap)
}
