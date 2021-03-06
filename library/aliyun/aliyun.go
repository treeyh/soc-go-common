package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/logger"
	"sync"
)

var (
	aLiYunProxys = make(map[string]*ALiYunProxy)

	_defaultConfigName = "default"

	poolMutex sync.Mutex

	log = logger.Logger()
)

type ALiYunProxy struct {
	aLiYunConfig *config.ALiYunConfig
	client       *sdk.Client
}

// InitWeChatConfig 初始化微信配置
func InitALiYunConfig(weChatConfigs map[string]config.ALiYunConfig) {
	poolMutex.Lock()
	defer poolMutex.Unlock()

	aLiYunProxys1 := make(map[string]*ALiYunProxy)
	for k, v := range weChatConfigs {
		client, err := sdk.NewClientWithAccessKey(v.RegionId, v.AccessKey, v.AccessKeySecret)
		if err != nil {
			log.Errorf(" init aliyun sdk error: %#v ", err)
			panic(err)
		}
		aLiYunProxys1[k] = &ALiYunProxy{aLiYunConfig: &v, client: client}
	}
	aLiYunProxys = aLiYunProxys1
}

// GetProxy get default redis oper proxy
func GetProxy() *ALiYunProxy {
	return GetProxyByName(_defaultConfigName)
}

// GetProxyByName get redis oper proxy
func GetProxyByName(name string) *ALiYunProxy {
	if v, ok := aLiYunProxys[name]; ok {
		return v
	}
	return nil
}
