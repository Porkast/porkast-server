package routers

import (
	"porkast-server/internal/ctls"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func V1ApiRouter(group *ghttp.RouterGroup) {
	group.Middleware(middleware.SetI18nLang)
	group.Middleware(middleware.MiddlewareCORS)
	unAuthGroup := group.Group("/")
	unAuthGroup.POST("/user/login", ctls.Ctl.DoLogin)
	unAuthGroup.POST("/user/register", ctls.Ctl.DoRegister)
	unAuthGroup.GET("/search/feed/item", ctls.Ctl.SearchFeedItemAPI)
	unAuthGroup.GET("/search/feed/channel", ctls.Ctl.SearchFeedChannelAPI)

	// search keyword subscription
	unAuthGroup.GET("/subscription/list/", ctls.Ctl.GetUserSubKeywordListAPI)
	unAuthGroup.GET("/subscription/:userId/:keyword", ctls.Ctl.GetSubKeywordItemListAPI)

	unAuthGroup.GET("/feed/channel/:id", ctls.Ctl.GetFeedChannelDetailAPI)
	unAuthGroup.GET("/feed/channel/:channelId/item/:itemId", ctls.Ctl.GetFeedItemDetailAPI)

	// auth group

	authGroup := group.Group("/")
	// listen later
	authGroup.Middleware(middleware.AuthToken)
	authGroup.POST("/listenlater/item", ctls.Ctl.AddListenLater)
	authGroup.GET("/listenlater/item", ctls.Ctl.GetListenLater)
	authGroup.GET("/listenlater/list", ctls.Ctl.GetListenLaterList)

	// search keyword subscription
	authGroup.POST("/subscription/keyword", ctls.Ctl.SubKeyword)

}
