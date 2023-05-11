package ctls

import (
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/service/middleware"

	feedService "guoshao-fm-web/internal/service/feed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) AddListenLater(req *ghttp.Request) {
	var (
		err     error
		reqData *AddListenLaterReqData
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	err = feedService.CreateListenLaterByUserIdAndFeedId(req.GetCtx(), reqData.UserId, reqData.ChannelId, reqData.ItemId)
	if err != nil {
		g.Log().Line().Error(req.GetCtx(), "add listen later failed :\n", err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}
	middleware.JsonExit(req, 0, "add listen later success", nil)

}

func (ctl *controller) GetListenLaterList(req *ghttp.Request) {

	var (
		err        error
		reqData    *GetListenLaterListReqData
		resultList []dto.UserListenLater
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	resultList, err = feedService.GetListenLaterListByUserId(req.GetCtx(), reqData.UserId)
	if err != nil {
		g.Log().Line().Error(req.GetCtx(), "get listen later failed :\n", err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}
	middleware.JsonExit(req, 0, "get listen later success", resultList)

}
