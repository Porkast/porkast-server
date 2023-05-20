package ctls

import (
	"fmt"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	feedService "guoshao-fm-web/internal/service/feed"
	"strconv"

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
	tplMap[consts.TOTAL_PAGE] = totalPage
	tplMap[consts.SEARCH_RESULT_COUNT_TEXT] = totalCountText
	tplMap[consts.SEARCH_TOOK_TIME_TEXT] = tookTimeText
	tplMap[consts.FEED_ITEMS] = items
	req.Response.WriteTpl("search.html", tplMap)
}
