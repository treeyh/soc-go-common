package wechat

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/utils/http_client"
	"github.com/treeyh/soc-go-common/core/utils/json"
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

func (wcp *WechatProxy) GetAccessToken(ctx context.Context) (string, errors.AppError) {

	if _accessToken == nil || _accessToken.ExpiresIn <= 0 || time.Now().Unix() > _accessToken.ExpiresTime.Unix() {
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

	if status != 200 || *result == "" {
		return &WechatAccessTokenResp{
			WechatBaseResp: WechatBaseResp{
				HttpStatus: status,
			},
		}, errors.NewAppError(errors.WechatRequestFail)
	}

	resp := &WechatAccessTokenResp{}
	err1 := json.FromJson(*result, resp)
	if err1 != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatRequestFail, err)
	}

	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}
	resp.HttpStatus = status

	return resp, nil
}
