package ctls

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/service/feed"
	"guoshao-fm-web/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) SubKeyword(req *ghttp.Request) {
	var (
		err     error
		ctx     context.Context
		reqData *SubSearchKeywordReqData
	)

	ctx = req.Context()
	err = req.Parse(&reqData)
	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	err = feed.SubFeedByKeyword(ctx, reqData.UserId, reqData.Keyword, reqData.Lang, reqData.SortByDate)
	if err != nil {
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(ctx, `{#sub_keyword_exist}`), nil)
		}
		g.Log().Line().Error(ctx, err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}

	middleware.JsonExit(req, 0, g.I18n().T(ctx, `{#sub_keyword_success}`), nil)
}
