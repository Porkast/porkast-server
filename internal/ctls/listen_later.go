package ctls

import (
	"fmt"
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/service/middleware"
	userService "porkast-server/internal/service/user"

	feedService "porkast-server/internal/service/feed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

func (ctl *controller) ListenLaterTpl(req *ghttp.Request) {
	var (
		ctx                        = gctx.New()
		err                        error
		tplMap                     = consts.GetCommonTplMap(req.GetCtx())
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
		if userListenLaterItemDtoList[0].Count%10 > 0 {
			totalPage = totalPage + 1
		}
	}
	tplMap[consts.LISTEN_LATER_ITEM_LIST] = userListenLaterItemDtoList
	tplMap[consts.USER_INFO] = userInfo
	tplMap[consts.CURRENT_PAGE] = page
	tplMap[consts.TOTAL_PAGE] = totalPage
	tplMap[consts.USER_LISTEN_LATER_TOTAL_COUNT] = gconv.String(totalCount) + g.I18n().T(ctx, `{#total_channe_items_count}`)
	tplMap[consts.LISTEN_LATER_PLAYLIST_NAME] = g.I18n().T(ctx, `{#play_list}`)
	tplMap[consts.LISTEN_LATER_PLAYLIST_TITLE] = g.I18n().Tf(ctx, "listen_later_rss_channel_title", userInfo.Nickname)
	tplMap[consts.LISTEN_LATER_PLAYLIST_RSS_LINK] = fmt.Sprintf("https://www.guoshaofm.com/listenlater/%s/rss", userInfo.Id)
	tplMap[consts.LISTEN_LATER_PLAYLIST_DESCRIPTION] = g.I18n().Tf(ctx, "listen_later_rss_channel_description", userInfo.Nickname)
	tplMap[consts.SUBSCRIPTION] = g.I18n().T(ctx, `{#subscription}`)
	tplMap[consts.COPY] = g.I18n().T(ctx, `{#copy}`)
	tplMap[consts.LISTEN_LATER_PLAYLIST_COPYRIGHT] = fmt.Sprintf("@%s GuoshaoFM", userInfo.Nickname)
	tplMap[consts.LISTEN_LATER_SUB_LIST] = g.I18n().T(ctx, `{#listen_later_sub_list}`)
	tplMap[consts.LISTEN_LATER_COPY_RSS_LINK] = g.I18n().T(ctx, `{#listen_later_copy_rss_link}`)
	tplMap[consts.LISTEN_LATER_COPY_TO_RSS_APP] = g.I18n().T(ctx, `{#listen_later_copy_to_rss_app}`)
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

	err = feedService.CreateListenLaterByUserIdAndFeedId(req.GetCtx(), reqData.UserId, reqData.ChannelId, reqData.ItemId, reqData.Source)
	if err != nil {
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(req.GetCtx(), `{#listen_later_exist}`), nil)
		} else {
			g.Log().Line().Error(req.GetCtx(), err)
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
