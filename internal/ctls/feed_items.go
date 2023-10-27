package ctls

import (
	"porkast-server/internal/dto"
	feedService "porkast-server/internal/service/feed"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)


func (ctl *controller) GetFeedItemDetailAPI(req *ghttp.Request) {
	var (
		err         error
		itemInfo    dto.FeedItem
		channelInfo dto.FeedChannel
		itemId      string
		channelId   string
	)

	itemId = req.Get("itemId").String()
	channelId = req.Get("channelId").String()
	channelInfo, itemInfo, err = feedService.GetFeedItemByItemId(req.Context(), channelId, itemId)
	if err != nil {
		middleware.JsonExit(req, 500, err.Error())
	}

	middleware.JsonExit(req, 200, "", g.Map{
		"itemInfo":    itemInfo,
		"channelInfo": channelInfo,
	})

}
