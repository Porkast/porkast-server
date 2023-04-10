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
		err           error
		searchKeyword string
		offset         int
		totalPage     int
		items         []dto.FeedItem
	)

	searchKeyword = req.GetQuery("q").String()
	offset = req.GetQuery("offset").Int()
	items, err = feedService.SearchFeedItemsByKeyword(req.Context(), searchKeyword, offset, 10)
	if err != nil {
		//TODO Add error page
	}
	req.Response.WriteTpl("search.html", g.Map{
		"searchKeyword":   searchKeyword,
		"currentPage":     offset,
		"totalPage":       totalPage,
		consts.FEED_ITEMS: items,
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
