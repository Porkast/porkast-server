package ctls

import (
	"guoshao-fm-web/internal/consts"
	feedService "guoshao-fm-web/internal/service/feed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) IndexTpl(req *ghttp.Request) {
	var (
		err               error
		tplMap            = consts.GetCommonTplMap()
		channelTotalCount int
		itemTotalCount    int
	)

	channelTotalCount, err = feedService.GetAllFeedChannelCountFromCache(req.GetCtx())
	if err != nil {
		// TODO: redirect to error page
	}

	itemTotalCount, err = feedService.GetAllFeedItemCountFromCache(req.GetCtx())
	if err != nil {
		// TODO: redirect to error page
	}

	tplMap[consts.TOTAL_CHANNEL_COUNT] = channelTotalCount
	tplMap[consts.TOTAL_ITEM_COUNT] = itemTotalCount
	if itemTotalCount != 0 && channelTotalCount != 0 {
		tplMap[consts.SEARCH_CN_FEED_ITEM_CHANNEL_TOTAL_WITH_COUNT] = g.I18n().Tf(req.GetCtx(), consts.SEARCH_CN_FEED_ITEM_CHANNEL_TOTAL_WITH_COUNT, itemTotalCount, channelTotalCount)
	} else {
		tplMap[consts.SEARCH_CN_FEED_ITEM_CHANNEL_TOTAL_WITH_COUNT] = g.I18n().Tf(req.GetCtx(), consts.SEARCH_PODCAST_BY_KEYWORD)
	}
	req.Response.WriteTpl("index.html", tplMap)
	return
}
