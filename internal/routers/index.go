package routers

import (
	"guoshao-fm-web/internal/ctls"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.GET("/", ctls.Ctl.IndexTpl)
	group.GET("/s/:keyword/:start", ctls.Ctl.SearchResult)
}
