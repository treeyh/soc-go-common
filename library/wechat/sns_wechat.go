package wechat

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/utils/http_client"
	"github.com/treeyh/soc-go-common/core/utils/json"
)

func (wcp *WechatProxy) JsCode2Session(ctx context.Context, jsCode string) (*WechatCode2SessionResp, errors.AppError) {
	url := wcp.GetPreUrl() + "/sns/jscode2session"
	params := make(map[string]string)
	params["appid"] = wcp.wechatConfig.AppId
	params["secret"] = wcp.wechatConfig.AppSecret
	params["js_code"] = jsCode
	params["grant_type"] = "authorization_code"

	result, status, err := http_client.Get(ctx, url, params)
	if err != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatRequestFail, err)
	}

	if status != 200 || result == "" {
		return &WechatCode2SessionResp{
			WechatErrorResp: WechatErrorResp{},
		}, errors.NewAppError(errors.WechatRequestFail)
	}

	resp := &WechatCode2SessionResp{}
	err1 := json.FromJson(result, resp)
	if err1 != nil {
		return nil, errors.NewAppErrorByExistError(errors.WechatRequestFail, err)
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil

}
