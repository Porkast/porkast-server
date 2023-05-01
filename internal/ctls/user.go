package ctls

import (
	"guoshao-fm-web/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) LoginTpl(req *ghttp.Request) {

	req.Response.WriteTpl("user/login.html", g.Map{
		consts.APP_NAME_KEY: consts.APP_NAME,
	})
}
