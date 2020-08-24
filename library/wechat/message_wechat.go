package wechat

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
)

func (wcp *WechatProxy) SendTemplateMessage(ctx context.Context, message *WechatTemplateMessageReq) (*WechatMessageResp, errors.AppError) {

	url := wcp.GetPreUrl() + "/cgi-bin/message/template/send"
	params := make(map[string]string)

	resp := &WechatMessageResp{}

	err := wcp.postJson(ctx, url, params, message, resp, true)

	if err != nil {
		return nil, err
	}
	if resp.ErrCode > 0 {
		return nil, errors.NewAppError(errors.WechatRequestError, resp.ErrCode, resp.ErrMsg)
	}

	return resp, nil
}
