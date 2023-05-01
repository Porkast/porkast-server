package routers

import (
	"guoshao-fm-web/internal/ctls"

	"github.com/gogf/gf/v2/net/ghttp"
)

func WebRouter(group *ghttp.RouterGroup) {
	group.GET("/", ctls.Ctl.IndexTpl)
	group.GET("/search", ctls.Ctl.SearchResult)
	group.GET("/feed/channel/:id", ctls.Ctl.FeedChannelDetail)
	group.GET("/feed/:channelId/item/:itemId", ctls.Ctl.FeedItemDetail)

	group.GET("/login", ctls.Ctl.LoginTpl)
}
