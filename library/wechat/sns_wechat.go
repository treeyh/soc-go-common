package wechat

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
)

func (wcp *WechatProxy) JsCode2Session(ctx context.Context, jsCode string) (*WechatCode2SessionResp, errors.AppError) {
	url := wcp.GetPreUrl() + "/sns/jscode2session"
	params := make(map[string]string)
	params["appid"] = wcp.wechatConfig.AppId
	params["secret"] = wcp.wechatConfig.AppSecret
	params["js_code"] = jsCode
	params["grant_type"] = "authorization_code"

	resp := &WechatCode2SessionResp{}

	err := wcp.getJson(ctx, url, params, resp, false)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}
