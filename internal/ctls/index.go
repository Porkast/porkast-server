package ctls

import (
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	feedService "guoshao-fm-web/internal/service/feed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) IndexTpl(req *ghttp.Request) {
	req.Response.WriteTpl("index.html", g.Map{
		"name": "锅烧FM",
	})
	return
}

func (ctl *controller) SearchResult(req *ghttp.Request) {
	var (
		searchKeyword string
		start         int
		totalPage     int
		feedItemList  []string
	)

	searchKeyword = req.Get("keyword").String()
	start = req.Get("start").Int()
	req.Response.WriteTpl("search.html", g.Map{
		"searchKeyword":   searchKeyword,
		"currentPage":     start,
		"totalPage":       totalPage,
		consts.FEED_ITEMS: feedItemList,
	})
}

func (ctl *controller) FeedChannelDetail(req *ghttp.Request) {
	var (
		err         error
		channelInfo dto.FeedChannel
		channelId   string
	)

	channelId = req.Get("id").String()
	channelInfo, err = feedService.GetChannelInfoByChannelId(req.Context(), channelId)
	if err != nil {
		// TODO redirect to error page
	}
	req.Response.WriteTpl("channel.html", g.Map{
		"channelInfo":     channelInfo,
		consts.FEED_ITEMS: channelInfo.Items,
	})
}

func (ctl *controller) FeedItemDetail(req *ghttp.Request) {
	var (
		err      error
		itemInfo dto.FeedItem
		itemId   string
	)

	itemId = req.Get("id").String()
	itemInfo, err = feedService.GetFeedItemByItemId(req.Context(), itemId)
	if err != nil {
		// TODO redirect to error page
	}
	req.Response.WriteTpl("item.html", g.Map{
		"itemInfo": itemInfo,
	})
}
