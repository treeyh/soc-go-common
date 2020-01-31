package wechat

import "github.com/treeyh/soc-go-common/core/utils/http"

func (wcp *WechatProxy) Jscode2session() {
	url := wcp.GetPreUrl() + "/sns/jscode2session"
	http.Get()

}
