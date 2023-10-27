package ctls

import (
	"porkast-server/internal/dto"
	feedService "porkast-server/internal/service/feed"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

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

	middleware.JsonExit(req, 0, "", g.Map{
		"channelInfo": channelInfo,
		"page": page,
	})
	
}
