package ctls

import (
	"guoshao-fm-web/internal/consts"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) IndexTpl(req *ghttp.Request) {
	req.Response.WriteTpl("index.html", consts.GetCommonTplMap())
	return
}

