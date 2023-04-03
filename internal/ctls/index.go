package ctls

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)


func (ctl *controller) IndexTpl(req *ghttp.Request) {
	req.Response.WriteTpl("index.html", g.Map{
		"name": "锅烧FM",
	})
	return
}
