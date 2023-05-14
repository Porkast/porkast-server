package ctls

import (
	"fmt"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	feedService "guoshao-fm-web/internal/service/feed"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) IndexTpl(req *ghttp.Request) {
	req.Response.WriteTpl("index.html", consts.GetCommonTplMap())
	return
}

func (ctl *controller) SearchResult(req *ghttp.Request) {
	var (
		err            error
		searchKeyword  string
		page           int
		totalPage      int
		totalCount     int
		totalCountText string
		tookTime       float64
		tookTimeStr    string
		tookTimeText   string
		items          []dto.FeedItem
	)

	searchKeyword = req.GetQuery("q").String()
	page = req.GetQuery("page").Int()
	items, err = feedService.SearchFeedItemsByKeyword(req.Context(), searchKeyword, page, 10)
	if err != nil {
		//TODO: Add error page
	}
	if len(items) > 0 {
		tookTime = items[0].TookTime
		tookTimeStr = strconv.FormatFloat(tookTime, 'f', -3, 64)
		totalCount = items[0].Count
		totalPage = totalCount / 10
	}
	totalCountText = fmt.Sprintf(consts.SEARCH_RESULT_COUNT_TEXT_VALUE, totalCount)
	tookTimeText = fmt.Sprintf(consts.SEARCH_TOOK_TIME_TEXT_VALUE, tookTimeStr)
	var tplMap = consts.GetCommonTplMap()
	tplMap[consts.SEARCH_KEY_WORD] = searchKeyword
	tplMap[consts.CURRENT_PAGE] = page
	tplMap[consts.SEARCH_KEY_WORD] = searchKeyword
	tplMap[consts.SEARCH_KEY_WORD] = searchKeyword
	tplMap[consts.SEARCH_KEY_WORD] = searchKeyword
	tplMap[consts.TOTAL_PAGE] = totalPage
	tplMap[consts.SEARCH_RESULT_COUNT_TEXT] = totalCountText
	tplMap[consts.SEARCH_TOOK_TIME_TEXT] = tookTimeText
	tplMap[consts.FEED_ITEMS] = items
	req.Response.WriteTpl("search.html", tplMap)
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
		// TODO: redirect to error page
	}
	var tplMap = consts.GetCommonTplMap()
	tplMap[consts.CHANNEL_INFO] = channelInfo
	tplMap[consts.FEED_ITEMS] = channelInfo.Items
	req.Response.WriteTpl("feed_channel.html", tplMap)
}

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
	var tplMap = consts.GetCommonTplMap()
	tplMap[consts.ITEM_INFO] = itemInfo
	tplMap[consts.CHANNEL_INFO] = channelInfo
	req.Response.WriteTpl("feed_item.html", g.Map{})
}
