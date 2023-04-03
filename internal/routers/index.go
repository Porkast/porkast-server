package routers

import (
	"guoshao-fm-web/internal/controllers"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.GET("/", controllers.Ctl.IndexTpl)
}
