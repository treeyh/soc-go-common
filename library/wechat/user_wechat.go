package wechat

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
)

func (wcp *WechatProxy) GetUserInfo(ctx context.Context, openid string) (*WechatUserInfo, errors.AppError) {

	url := wcp.GetPreUrl() + "/cgi-bin/user/info"
	params := make(map[string]string)
	params["openid"] = openid
	params["lang"] = ""

	resp := &WechatUserInfo{}

	err := wcp.getJson(ctx, url, params, resp, true)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}

func (wcp *WechatProxy) GetUserInfoBatch(ctx context.Context, userInfoReqs []WechatUserInfoReq) (*WechatUserInfoList, errors.AppError) {

	url := wcp.GetPreUrl() + "/cgi-bin/user/info/batchget"
	params := make(map[string]string)

	data := make(map[string][]WechatUserInfoReq)
	for _, v := range userInfoReqs {
		if v.Lang == "" {
			v.Lang = WechatLang_ZH_CN
		}
	}
	data["user_list"] = userInfoReqs

	resp := &WechatUserInfoList{}

	err := wcp.postJson(ctx, url, params, data, resp, true)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}

func (wcp *WechatProxy) GetOpenids(ctx context.Context, nextOpenid string) (*WechatOpenidsResp, errors.AppError) {

	url := wcp.GetPreUrl() + "/cgi-bin/user/get"
	params := make(map[string]string)
	params["next_openid"] = nextOpenid

	resp := &WechatOpenidsResp{}
	err := wcp.getJson(ctx, url, params, resp, true)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}
