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

func (wcp *WechatProxy) GetAccessTokenByCode(ctx context.Context, code string) (*WechatCode2AccessTokenResp, errors.AppError) {
	url := wcp.GetPreUrl() + "/sns/oauth2/access_token"
	params := make(map[string]string)
	params["appid"] = wcp.wechatConfig.AppId
	params["secret"] = wcp.wechatConfig.AppSecret
	params["code"] = code
	params["grant_type"] = "authorization_code"

	resp := &WechatCode2AccessTokenResp{}

	err := wcp.getJson(ctx, url, params, resp, false)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}

func (wcp *WechatProxy) RefreshToken(ctx context.Context, refreshToken string) (*WechatCode2AccessTokenResp, errors.AppError) {
	url := wcp.GetPreUrl() + "/sns/oauth2/refresh_token"
	params := make(map[string]string)
	params["appid"] = wcp.wechatConfig.AppId
	params["refresh_token"] = refreshToken
	params["grant_type"] = "refresh_token"

	resp := &WechatCode2AccessTokenResp{}

	err := wcp.getJson(ctx, url, params, resp, false)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}

func (wcp *WechatProxy) Auth(ctx context.Context, accessToken, openid string) (*WechatErrorResp, errors.AppError) {
	url := wcp.GetPreUrl() + "/sns/auth"
	params := make(map[string]string)
	params["access_token"] = accessToken
	params["openid"] = openid

	resp := &WechatErrorResp{}

	err := wcp.getJson(ctx, url, params, resp, false)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}

func (wcp *WechatProxy) GetSnsUserInfo(ctx context.Context, accessToken, openid, lang string) (*WechatUserInfo, errors.AppError) {
	url := wcp.GetPreUrl() + "/sns/userinfo"
	params := make(map[string]string)
	params["access_token"] = accessToken
	params["openid"] = openid
	params["lang"] = lang

	resp := &WechatUserInfo{}

	err := wcp.getJson(ctx, url, params, resp, false)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}
