package wechat

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/utils/http_client"
	"github.com/treeyh/soc-go-common/core/utils/json"
)

func (wcp *WechatProxy) GetAccessToken(ctx context.Context) (*WechatAccessTokenResp, errors.AppError) {
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
