package routers

import (
	"guoshao-fm-web/internal/ctls"
	"guoshao-fm-web/internal/service/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.GET("/", ctls.Ctl.IndexTpl)
	group.GET("/search", ctls.Ctl.SearchResult)
	group.GET("/feed/channel/:id", ctls.Ctl.FeedChannelDetail)
	group.GET("/feed/:channelId/item/:itemId", ctls.Ctl.FeedItemDetail)

	group.GET("/login", ctls.Ctl.LoginTpl)
	group.GET("/register", ctls.Ctl.RegisterTpl)

}

func V1ApiRouter(group *ghttp.RouterGroup) {
	unAuthGroup := group.Group("/")
	unAuthGroup.POST("/user/login", ctls.Ctl.DoLogin)
	unAuthGroup.POST("/user/register", ctls.Ctl.DoRegister)

	authGroup := group.Group("/")
	authGroup.Middleware(middleware.AuthToken)
	authGroup.POST("/listenlater/item", ctls.Ctl.AddListenLater)
	authGroup.GET("/listenlater/list", ctls.Ctl.GetListenLaterList)
}
