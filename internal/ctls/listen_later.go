package ctls

import (
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/service/middleware"

	feedService "porkast-server/internal/service/feed"

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

	err = feedService.CreateListenLaterByUserIdAndFeedId(req.GetCtx(), reqData.UserId, reqData.ChannelId, reqData.ItemId, reqData.Source)
	if err != nil {
		g.Log().Line().Error(req.GetCtx(), err)
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(req.GetCtx(), `{#listen_later_exist}`), nil)
		} else {
			middleware.JsonExit(req, 1, err.Error(), nil)
		}
	}
	middleware.JsonExit(req, 0, g.I18n().T(req.GetCtx(), `{#add_listen_later_sucess}`), nil)

}

func (ctl *controller) GetListenLater(req *ghttp.Request) {
	var (
		err                error
		reqData            *GetListenLaterReqData
		userListenLaterDto dto.UserListenLater
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	userListenLaterDto, err = feedService.GetListenLaterByUserIdAndFeedId(req.GetCtx(), reqData.UserId, reqData.ChannelId, reqData.ItemId)
	if err != nil {
		g.Log().Line().Error(req.GetCtx(), "get listen later failed :\n", err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}
	middleware.JsonExit(req, 0, "get listen later success", userListenLaterDto)

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

	resultList, err = feedService.GetListenLaterListByUserId(req.GetCtx(), reqData.UserId, reqData.Offset, reqData.Limit)
	if err != nil {
		g.Log().Line().Error(req.GetCtx(), "get listen later failed :\n", err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}
	middleware.JsonExit(req, 0, "get listen later success", resultList)

}

func (ctl *controller) GetListenLaterRSS(req *ghttp.Request) {
	var (
		err    error
		userId string
		rss    string
	)

	userId = req.Get("userId").String()

	rss, err = feedService.GetListenLaterRSSByUserId(req.GetCtx(), userId)
	if err != nil {
		// TODO: redirect to error page
	}

	req.Response.WriteXml(rss)
}
