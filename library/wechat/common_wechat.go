package wechat

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/utils/http_client"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var (
	_lock        sync.Mutex
	_accessToken *WechatAccessToken
)

type WechatAccessToken struct {
	AccessToken string
	ExpiresIn   int
	ExpiresTime time.Time
}

// GetAccessToken 获取accesstoken，isForce是否强制获取
func (wcp *WechatProxy) GetAccessToken(ctx context.Context, isForce bool) (string, errors.AppError) {

	if _accessToken == nil || isForce || _accessToken.ExpiresIn <= 0 || time.Now().Unix() > _accessToken.ExpiresTime.Unix() {
		_lock.Lock()
		defer _lock.Unlock()
		// 重复判断，防止请求两次
		// TODO 由于是单服务，未考虑分布式使用缓存
		if _accessToken == nil || _accessToken.ExpiresIn <= 0 || time.Now().Unix() > _accessToken.ExpiresTime.Unix() {
			resp, err := wcp.getAccessToken(ctx)
			if err != nil {
				return "", err
			}
			_accessToken = &WechatAccessToken{
				AccessToken: resp.AccessToken,
				ExpiresIn:   resp.ExpiresIn,
				// 提前10分钟到期
				ExpiresTime: time.Now().Add(time.Duration(resp.ExpiresIn-600) * time.Second),
			}
		}
	}
	return _accessToken.AccessToken, nil
}

func (wcp *WechatProxy) getAccessToken(ctx context.Context) (*WechatAccessTokenResp, errors.AppError) {
	url := wcp.GetPreUrl() + "/cgi-bin/token"

	params := make(map[string]string)
	params["appid"] = wcp.wechatConfig.AppId
	params["secret"] = wcp.wechatConfig.AppSecret
	params["grant_type"] = "client_credential"

	result, status, err := http_client.Get(ctx, url, params)
	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatRequestFail, err)
	}

	if status != 200 || result == "" {
		return &WechatAccessTokenResp{
			WechatErrorResp: WechatErrorResp{},
		}, errors.NewAppError(errors.WechatRequestFail)
	}

	resp := &WechatAccessTokenResp{}
	err1 := json.FromJson(result, resp)
	if err1 != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatRequestFail, err)
	}

	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}

// getJson 请求微信接口获取json返回
func (wcp *WechatProxy) getJson(ctx context.Context, url string, params map[string]string, resp interface{}) errors.AppError {

	ErrorStructValue, ErrorErrCodeValue := checkResponse(resp)

	token, err := wcp.GetAccessToken(ctx, false)
	if err != nil {
		return err
	}

	hasRetried := false
RETRY:

	str, httpStatus, err := http_client.Get(ctx, url, params)
	if err != nil {
		return err
	}
	if httpStatus != http.StatusOK {
		return errors.NewAppError(errors.WechatRequestFail)
	}
	err1 := json.FromJson(str, resp)
	if err1 != nil {
		return errors.NewAppErrorByExistError(errors.JsonDecodeFail, err1)
	}

	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return nil
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		log.ErrorCtx(ctx, "code:"+strconv.FormatInt(errCode, 10)+";msg:"+errMsg+";token:"+token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)

			token, err = wcp.GetAccessToken(ctx, true)
			if err != nil {
				return err
			}
			goto RETRY
		}
		fallthrough
	default:
		return nil
	}
}

// postJson 请求微信接口获取json返回
func (wcp *WechatProxy) postJson(ctx context.Context, url string, params map[string]string, data interface{}, resp interface{}) errors.AppError {

	ErrorStructValue, ErrorErrCodeValue := checkResponse(resp)

	token, err := wcp.GetAccessToken(ctx, false)
	if err != nil {
		return err
	}

	body, err1 := json.ToJson(data)
	if err1 != nil {
		return errors.NewAppErrorByExistError(errors.JsonEncodeFail, err1)
	}

	hasRetried := false
RETRY:

	str, httpStatus, err := http_client.Post(ctx, url, params, body)
	if err != nil {
		return err
	}
	if httpStatus != http.StatusOK {
		return errors.NewAppError(errors.WechatRequestFail)
	}
	err1 = json.FromJson(str, resp)
	if err1 != nil {
		return errors.NewAppErrorByExistError(errors.JsonDecodeFail, err1)
	}

	switch errCode := ErrorErrCodeValue.Int(); errCode {
	case ErrCodeOK:
		return nil
	case ErrCodeInvalidCredential, ErrCodeAccessTokenExpired:
		errMsg := ErrorStructValue.Field(errorErrMsgIndex).String()
		log.ErrorCtx(ctx, "code:"+strconv.FormatInt(errCode, 10)+";msg:"+errMsg+";token:"+token)
		if !hasRetried {
			hasRetried = true
			ErrorStructValue.Set(errorZeroValue)

			token, err = wcp.GetAccessToken(ctx, true)
			if err != nil {
				return err
			}
			goto RETRY
		}
		fallthrough
	default:
		return nil
	}
}

// checkResponse 检查 response 参数是否满足特定的结构要求, 如果不满足要求则会 panic, 否则返回相应的 reflect.Value.
func checkResponse(response interface{}) (ErrorStructValue, ErrorErrCodeValue reflect.Value) {
	responseValue := reflect.ValueOf(response)
	if responseValue.Kind() != reflect.Ptr {
		panic("the type of response is incorrect")
	}
	responseStructValue := responseValue.Elem()
	if responseStructValue.Kind() != reflect.Struct {
		panic("the type of response is incorrect")
	}

	if t := responseStructValue.Type(); t == errorType {
		ErrorStructValue = responseStructValue
	} else {
		if t.NumField() == 0 {
			panic("the type of response is incorrect")
		}
		v := responseStructValue.Field(0)
		if v.Type() != errorType {
			panic("the type of response is incorrect")
		}
		ErrorStructValue = v
	}
	ErrorErrCodeValue = ErrorStructValue.Field(errorErrCodeIndex)
	return
}
