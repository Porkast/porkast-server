package ctls

import (
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/service/middleware"

	feedService "guoshao-fm-web/internal/service/feed"
	userService "guoshao-fm-web/internal/service/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

func (ctl *controller) ListenLaterTpl(req *ghttp.Request) {
	var (
		ctx                        = gctx.New()
		err                        error
		tplMap                     = consts.GetCommonTplMap()
		userId                     string
		userInfo                   dto.UserInfo
		userListenLaterItemDtoList []dto.UserListenLater
	)

	userId = req.Get("userId").String()
	userInfo, err = userService.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		// TODO: redirect to error page
	}
	userListenLaterItemDtoList, err = feedService.GetListenLaterListByUserId(ctx, userId, 0, 10)
	if err != nil {
		// TODO: redirect to error page
	}
	tplMap[consts.LISTEN_LATER_ITEM_LIST] = userListenLaterItemDtoList
	tplMap[consts.USER_INFO] = userInfo
	req.Response.WriteTpl("listen_later_playlist.html", tplMap)
}

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
