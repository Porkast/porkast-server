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
	req.Response.WriteTpl("index.html", g.Map{
		consts.APP_NAME_KEY: consts.APP_NAME,
	})
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
	req.Response.WriteTpl("search.html", g.Map{
		consts.APP_NAME_KEY:             consts.APP_NAME,
		consts.SEARCH_KEY_WORD:          searchKeyword,
		consts.CURRENT_PAGE:             page,
		consts.TOTAL_PAGE:               totalPage,
		consts.SEARCH_RESULT_COUNT_TEXT: totalCountText,
		consts.SEARCH_TOOK_TIME_TEXT:    tookTimeText,
		consts.FEED_ITEMS:               items,
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
		// TODO: redirect to error page
	}
	req.Response.WriteTpl("feed_channel.html", g.Map{
		consts.CHANNEL_INFO: channelInfo,
		consts.FEED_ITEMS:   channelInfo.Items,
	})
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
	req.Response.WriteTpl("feed_item.html", g.Map{
		consts.ITEM_INFO:    itemInfo,
		consts.CHANNEL_INFO: channelInfo,
	})
}
