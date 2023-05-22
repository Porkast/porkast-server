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
	"github.com/gogf/gf/v2/util/gconv"
)

func (ctl *controller) ListenLaterTpl(req *ghttp.Request) {
	var (
		ctx                        = gctx.New()
		err                        error
		tplMap                     = consts.GetCommonTplMap()
		userId                     string
		userInfo                   dto.UserInfo
		totalCount                 int
		page                       int
		offset                     int
		totalPage                  int
		userListenLaterItemDtoList []dto.UserListenLater
	)

	userId = req.Get("userId").String()
	page = req.GetQuery("page").Int()
	if page == 0 {
		page = 1
		offset = 0
	} else {
		offset = (page - 1) * 10
	}
	userInfo, err = userService.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		// TODO: redirect to error page
	}
	userListenLaterItemDtoList, err = feedService.GetListenLaterListByUserId(ctx, userId, offset, 10)
	if err != nil {
		// TODO: redirect to error page
	}
	if len(userListenLaterItemDtoList) > 0 {
		totalCount = userListenLaterItemDtoList[0].Count
		totalPage = totalCount / 10
	}
	tplMap[consts.LISTEN_LATER_ITEM_LIST] = userListenLaterItemDtoList
	tplMap[consts.USER_INFO] = userInfo
	tplMap[consts.CURRENT_PAGE] = page
	tplMap[consts.TOTAL_PAGE] = totalPage
	tplMap[consts.USER_LISTEN_LATER_TOTAL_COUNT] = gconv.String(totalCount) + g.I18n().T(ctx, `{#total_channe_items_count}`)
	tplMap[consts.LISTEN_LATER_PLAYLIST_NAME] = g.I18n().T(ctx, `{#play_list}`)
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
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(req.GetCtx(), `{#listen_later_exist}`), nil)
		}
		g.Log().Line().Error(req.GetCtx(), err)
		middleware.JsonExit(req, 1, err.Error(), nil)
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
