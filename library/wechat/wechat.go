package wechat

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/treeyh/soc-go-common/core/logger"
	"strings"
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
func InitWeChatConfig(weChatConfigs map[string]config.WeChatConfig) {
	poolMutex.Lock()
	defer poolMutex.Unlock()

	wechatProxys1 := make(map[string]*WechatProxy)
	for k, v := range weChatConfigs {
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
func (wcp *WechatProxy) DecryptEncryptedData(ctx context.Context, sessionKey string, encryptedData string, iv string, isJSON bool) (interface{}, errors.AppError) {
	if len(sessionKey) != 24 {
		return nil, errors.NewAppError(errors.WechatOperationError, "sessionKey length is error")
	}

	aesCipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		log.Error("encryptedData decode base64 error :"+encryptedData+" error:"+err.Error(), logger.GetTraceField(ctx))
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "decode base64 error")
	}

	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		log.Error("sessionKey decode base64 error :"+sessionKey+" error:"+err.Error(), logger.GetTraceField(ctx))
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "decode base64 error")
	}

	if len(iv) != 24 {
		log.Error("iv length is error :"+iv, logger.GetTraceField(ctx))
		return nil, errors.NewAppError(errors.WechatOperationError, "iv length is error")
	}
	aesIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		log.Error("iv decode base64 error :"+iv+" error:"+err.Error(), logger.GetTraceField(ctx))
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "decode base64 error")
	}
	//aesPlantText := make([]byte, len(aesCipherText))

	dataBytes, err1 := AesDecrypt(ctx, aesCipherText, aesKey, aesIV)
	if err1 != nil {
		return nil, err1
	}
	jsonStr := string(dataBytes)
	log.InfoCtx(ctx, jsonStr)

	var decrypted map[string]interface{}
	err = json.Unmarshal(dataBytes, &decrypted)
	if err != nil {
		log.Error("format json error:"+jsonStr+" error:"+err.Error(), logger.GetTraceField(ctx))
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "format json error")
	}

	if decrypted["watermark"].(map[string]interface{})["appid"] != wcp.wechatConfig.AppId {
		log.Error("appID is not match:", logger.GetTraceField(ctx))
		return nil, errors.NewAppError(errors.WechatOperationError, "appID is not match")
	}

	if isJSON == true {
		return strings.TrimSpace(jsonStr), nil
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

// CheckErrCodeSucceed 检查接口返回是否成功
func CheckErrCodeSucceed(httpStatus int, errCode int64) bool {
	return httpStatus == 200 && errCode == 0
}

func (wcp *WechatProxy) DecryptWXOpenData(ctx context.Context, sessionKey, encryptData, iv string) (map[string]interface{}, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	dataBytes, err := AesDecrypt(ctx, decodeBytes, sessionKeyBytes, ivBytes)
	fmt.Println(string(dataBytes))
	m := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	temp := m["watermark"].(map[string]interface{})
	appid := temp["appid"].(string)
	if appid != wcp.wechatConfig.AppId {
		return nil, fmt.Errorf("invalid appid, get !%s!", appid)
	}
	if err != nil {
		return nil, err
	}
	return m, nil

}

func AesDecrypt(ctx context.Context, crypted, key, iv []byte) ([]byte, errors.AppError) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.ErrorCtx(ctx, "aesKey aes error:"+err.Error())
		return nil, errors.NewAppErrorByExistError(errors.WechatOperationError, err, "aes error")
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x0e' || ch == '\x0f' {
			origData[i] = ' '
		}
	}
	//{"phoneNumber":"15082726017","purePhoneNumber":"15082726017","countryCode":"86","watermark":{"timestamp":1539657521,"appid":"wx4c6c3ed14736228c"}}//<nil>
	return origData, nil
}
