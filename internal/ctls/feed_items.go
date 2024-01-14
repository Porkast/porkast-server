package ctls

import (
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	feedService "porkast-server/internal/service/feed"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) FeedItemDetail(req *ghttp.Request) {
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
		// TODO: redirect to error page
	}
	var tplMap = consts.GetCommonTplMap(req.GetCtx())
	tplMap[consts.ITEM_INFO] = itemInfo
	tplMap[consts.CHANNEL_INFO] = channelInfo
	req.Response.WriteTpl("feed_item.html", tplMap)
}

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

func (ctl *controller) ShareFeedItemTpl(req *ghttp.Request) {
	var (
		err         error
		itemInfo    dto.FeedItem
		channelInfo dto.FeedChannel
		itemId      string
		channelId   string
	)
	itemId = req.Get("itemId").String()
	channelId = req.Get("channelId").String()
	_, itemInfo, err = feedService.GetFeedItemByItemId(req.Context(), channelId, itemId)
	if err != nil {
		// TODO: redirect to error page
	}

	channelInfo, err = feedService.GetChannelInfoByChannelId(req.Context(), channelId, 0, 10)
	if err != nil {
		// TODO: redirect to error page
	}

	var tplMap = consts.GetCommonTplMap(req.GetCtx())
	tplMap[consts.ITEM_INFO] = itemInfo
	tplMap[consts.CHANNEL_INFO] = channelInfo
	tplMap[consts.FEED_ITEMS] = channelInfo.Items
	tplMap[consts.PAST_FEED_ITEMS] = g.I18n().T(req.GetCtx(), `{#past_feed_item}`)
	req.Response.WriteTpl("share_feed_item.html", tplMap)
}
