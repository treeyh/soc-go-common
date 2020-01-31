package wechat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"github.com/treeyh/soc-go-common/core/logger"
	"regexp"
	"sync"

	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/errors"
)

const weChatPreUrl = "https://api.weixin.qq.com"

var (
	wechatProxys = make(map[string]*WechatProxy)

	_defaultConfigName = "default"

	poolMutex sync.Mutex

	log = logger.Logger()
)

type WechatProxy struct {
	wechatConfig *config.WeChatConfig
}

func (wcp *WechatProxy) GetPreUrl() string {
	if wcp.wechatConfig.Host != "" {
		return wcp.wechatConfig.Host
	}
	return weChatPreUrl
}

// InitWeChatConfig 初始化微信配置
func InitWeChatConfig(weChatConfigs *map[string]config.WeChatConfig) {
	poolMutex.Lock()
	defer poolMutex.Unlock()

	wechatProxys1 := make(map[string]*WechatProxy)
	for k, v := range *weChatConfigs {
		wechatProxys1[k] = &WechatProxy{wechatConfig: &v}
	}
	wechatProxys = wechatProxys1
}

// GetProxy get default redis oper proxy
func GetProxy() *WechatProxy {
	return GetProxyByName(_defaultConfigName)
}

// GetProxyByName get redis oper proxy
func GetProxyByName(name string) *WechatProxy {
	if v, ok := wechatProxys[name]; ok {
		return v
	}
	return nil
}

// Decrypt Weixin APP's AES Data
// If isJSON is true, Decrypt return JSON type.
// If isJSON is false, Decrypt return map type.
func Decrypt(appId string, sessionKey string, encryptedData string, iv string, isJSON bool) (interface{}, errors.AppError) {
	if len(sessionKey) != 24 {
		return nil, errors.NewAppError(errors.WechatOperationError, "sessionKey length is error")
	}
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, errors.NewAppError(errors.WechatOperationError, "decode base64 error")
	}

	if len(iv) != 24 {
		return nil, errors.NewAppError(errors.WechatOperationError, "iv length is error")
	}
	aesIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "decode base64 error")
	}

	aesCipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "decode base64 error")
	}
	aesPlantText := make([]byte, len(aesCipherText))

	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "aes error")
	}

	mode := cipher.NewCBCDecrypter(aesBlock, aesIV)
	mode.CryptBlocks(aesPlantText, aesCipherText)
	aesPlantText = PKCS7UnPadding(aesPlantText)

	var decrypted map[string]interface{}

	re := regexp.MustCompile(`[^\{]*(\{.*\})[^\}]*`)
	aesPlantText = []byte(re.ReplaceAllString(string(aesPlantText), "$1"))

	err = json.Unmarshal(aesPlantText, &decrypted)
	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "format json error")
	}

	if decrypted["watermark"].(map[string]interface{})["appid"] != appId {
		return nil, errors.NewAppError(errors.WechatOperationError, "appID is not match")
	}

	if isJSON == true {
		return string(aesPlantText), nil
	}

	return decrypted, nil
}

// PKCS7UnPadding return unpadding []Byte plantText
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	if length > 0 {
		unPadding := int(plantText[length-1])
		return plantText[:(length - unPadding)]
	}
	return plantText
}
