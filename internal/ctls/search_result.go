package ctls

import (
	"guoshao-fm-web/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

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
