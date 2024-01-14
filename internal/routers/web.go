package routers

import (
	"porkast-server/internal/ctls"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.Middleware(middleware.SetI18nLang)
	group.GET("/", ctls.Ctl.IndexTpl)
	group.GET("/search", ctls.Ctl.SearchResult)
	group.GET("/feed/channel/:id", ctls.Ctl.FeedChannelDetail)
	group.GET("/feed/:channelId/item/:itemId", ctls.Ctl.FeedItemDetail)
	group.GET("/share/feed/:channelId/item/:itemId", ctls.Ctl.ShareFeedItemTpl)
	group.GET("/share/feed/channel/:channelId", ctls.Ctl.ShareFeedChannelTpl)

	group.GET("/login", ctls.Ctl.LoginTpl)
	group.GET("/register", ctls.Ctl.RegisterTpl)
	group.GET("/user/info/:id", ctls.Ctl.UserInfoTpl)

	group.GET("/listenlater/playlist/:userId", ctls.Ctl.ListenLaterTpl)
	group.GET("/listenlater/:userId/rss", ctls.Ctl.GetListenLaterRSS)

	group.GET("/subscription/:userId/:keyword/rss", ctls.Ctl.GetSubKeywordFeedRSS)
	group.GET("/user/sub/list/:userId/:page", ctls.Ctl.UserSubListTplt)
}

func V1ApiRouter(group *ghttp.RouterGroup) {
	group.Middleware(middleware.SetI18nLang)
	group.Middleware(middleware.MiddlewareCORS)
	unAuthGroup := group.Group("/")
	unAuthGroup.POST("/user/login", ctls.Ctl.DoLogin)
	unAuthGroup.POST("/user/register", ctls.Ctl.DoRegister)
	unAuthGroup.GET("/user/info/:userId", ctls.Ctl.GetUserInfo)
	unAuthGroup.POST("/user/sync", ctls.Ctl.SyncUserInfo)
	unAuthGroup.GET("/search/feed/item", ctls.Ctl.SearchFeedItemAPI)
	unAuthGroup.GET("/search/feed/channel", ctls.Ctl.SearchFeedChannelAPI)

	// search keyword subscription
	unAuthGroup.GET("/subscription/list/", ctls.Ctl.GetUserSubKeywordListAPI)
	unAuthGroup.GET("/subscription/:userId/:keyword", ctls.Ctl.GetSubKeywordItemListAPI)

	unAuthGroup.GET("/feed/channel/:id", ctls.Ctl.GetFeedChannelDetailAPI)
	unAuthGroup.GET("/feed/channel/:channelId/item/:itemId", ctls.Ctl.GetFeedItemDetailAPI)

	// playlist
	unAuthGroup.GET("/playlist/list/:userId", ctls.Ctl.GetUserPlaylistsByUserId)
	unAuthGroup.GET("/playlist/list/:userId/:playlistId", ctls.Ctl.GetUserPlaylistItemList)
	unAuthGroup.GET("/playlist/:playlistId", ctls.Ctl.GetPlaylistInfo)

	// auth group

	authGroup := group.Group("/")
	// listen later
	authGroup.Middleware(middleware.AuthToken)
	authGroup.POST("/listenlater/item", ctls.Ctl.AddListenLater)
	authGroup.GET("/listenlater/item", ctls.Ctl.GetListenLater)
	authGroup.GET("/listenlater/list", ctls.Ctl.GetListenLaterList)

	// search keyword subscription
	authGroup.POST("/subscription/keyword", ctls.Ctl.SubKeyword)

	// playlist
	authGroup.POST("/playlist", ctls.Ctl.CreatePlaylist)
	authGroup.POST("/playlist/subcription", ctls.Ctl.SubscribePlaylist)
	authGroup.POST("/playlist/item", ctls.Ctl.AddFeedItemToPlaylist)
}
