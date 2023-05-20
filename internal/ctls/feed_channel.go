package ctls

import (
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	feedService "guoshao-fm-web/internal/service/feed"

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
	var tplMap = consts.GetCommonTplMap()
	tplMap[consts.CHANNEL_INFO] = channelInfo
	tplMap[consts.FEED_ITEMS] = channelInfo.Items
	tplMap[consts.CURRENT_PAGE] = page
	tplMap[consts.TOTAL_CHANNE_ITEMS_COUNT] = gconv.String(channelInfo.Count) + g.I18n().T(req.GetCtx(), `{#total_channe_items_count}`)
	req.Response.WriteTpl("feed_channel.html", tplMap)
}
