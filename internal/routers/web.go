package routers

import (
	"guoshao-fm-web/internal/ctls"
	"guoshao-fm-web/internal/service/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.Middleware(middleware.SetI18nLang)
	group.GET("/", ctls.Ctl.IndexTpl)
	group.GET("/search", ctls.Ctl.SearchResult)
	group.GET("/feed/channel/:id", ctls.Ctl.FeedChannelDetail)
	group.GET("/feed/:channelId/item/:itemId", ctls.Ctl.FeedItemDetail)
	group.GET("/share/feed/:channelId/item/:itemId", ctls.Ctl.ShareFeedItemTpl)
    group.GET("/share/feed/channel/:channelId", ctls.Ctl.ShareFeedItemTpl)

	group.GET("/login", ctls.Ctl.LoginTpl)
	group.GET("/register", ctls.Ctl.RegisterTpl)

	group.GET("/listenlater/playlist/:userId", ctls.Ctl.ListenLaterTpl)

	group.GET("/user/info/:id", ctls.Ctl.UserInfoTpl)
}

func V1ApiRouter(group *ghttp.RouterGroup) {
	group.Middleware(middleware.SetI18nLang)
	unAuthGroup := group.Group("/")
	unAuthGroup.POST("/user/login", ctls.Ctl.DoLogin)
	unAuthGroup.POST("/user/register", ctls.Ctl.DoRegister)

	authGroup := group.Group("/")
	authGroup.Middleware(middleware.AuthToken)
	authGroup.POST("/listenlater/item", ctls.Ctl.AddListenLater)
	authGroup.GET("/listenlater/item", ctls.Ctl.GetListenLater)
	authGroup.GET("/listenlater/list", ctls.Ctl.GetListenLaterList)
}
