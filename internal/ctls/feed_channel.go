package ctls

import (
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	feedService "porkast-server/internal/service/feed"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func (ctl *controller) FeedChannelDetail(req *ghttp.Request) {
	var (
		err         error
		channelInfo dto.FeedChannel
		channelId   string
		page        int
		offset      int
		limit       = 10
		totalPage   int
	)

	channelId = req.Get("id").String()
	page = req.GetQuery("page").Int()
	if page == 0 {
		page = 1
		offset = 0
	} else {
		offset = (page - 1) * 10
	}

	channelInfo, err = feedService.GetChannelInfoByChannelId(req.Context(), channelId, offset, limit)
	if err != nil {
		// TODO: redirect to error page
	}

	if len(channelInfo.Items) > 0 {
		totalPage = channelInfo.Count / 10
		if channelInfo.Count%10 > 0 {
			totalPage = totalPage + 1
		}
	}

	var tplMap = consts.GetCommonTplMap(req.GetCtx())
	tplMap[consts.CHANNEL_INFO] = channelInfo
	tplMap[consts.FEED_ITEMS] = channelInfo.Items
	tplMap[consts.CURRENT_PAGE] = page
	tplMap[consts.TOTAL_PAGE] = totalPage
	tplMap[consts.TOTAL_CHANNE_ITEMS_COUNT] = gconv.String(channelInfo.Count) + g.I18n().T(req.GetCtx(), `{#total_channe_items_count}`)
	req.Response.WriteTpl("feed_channel.html", tplMap)
}

func (ctl *controller) GetFeedChannelDetailAPI(req *ghttp.Request)  {
	var (
		err         error
		channelInfo dto.FeedChannel
		channelId   string
		page        int
		offset      int
		limit       = 10
	)

	channelId = req.Get("id").String()
	page = req.GetQuery("page").Int()
	if page == 0 {
		page = 1
		offset = 0
	} else {
		offset = (page - 1) * 10
	}

	channelInfo, err = feedService.GetChannelInfoByChannelId(req.Context(), channelId, offset, limit)
	if err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	middleware.JsonExit(req, 0, "", channelInfo)
	
}

func (ctl *controller) ShareFeedChannelTpl(req *ghttp.Request) {
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
	req.Response.WriteTpl("share_feed_channel.html", tplMap)
}
