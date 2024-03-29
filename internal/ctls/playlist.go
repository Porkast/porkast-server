package ctls

import (
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/service/feed"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) CreatePlaylist(req *ghttp.Request) {
	var (
		ctx     = req.GetCtx()
		err     error
		reqData *CreatePlaylistReqData
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	err = feed.CreatePlaylist(ctx, reqData.Name, reqData.UserId, reqData.Description)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(req.GetCtx(), `{#already_added_to_playlist}`), nil)
		} else {
			middleware.JsonExit(req, 1, err.Error(), nil)
		}
	}

	middleware.JsonExit(req, 0, g.I18n().T(ctx, `{#create_playlist_sucess}`), nil)

}

func (ctl *controller) GetPlaylistInfo(req *ghttp.Request) {
	var (
		ctx     = req.GetCtx()
		err     error
		reqData *GetPlaylistInfoReqData
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	playlistInfo, err := feed.GetPlaylistByPlaylistId(ctx, reqData.PlaylistId)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}

	middleware.JsonExit(req, 0, "", playlistInfo)

}

func (ctl *controller) SubscribePlaylist(req *ghttp.Request) {
	var (
		ctx     = req.GetCtx()
		err     error
		reqData *SubscribePlaylistReqData
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	err = feed.SubscribePlaylist(ctx, reqData.UserId, reqData.PlaylistId)

	if err != nil {
		g.Log().Line().Error(ctx, err)
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(req.GetCtx(), `{#already_added_to_playlist}`), nil)
		} else {
			middleware.JsonExit(req, 1, err.Error(), nil)
		}
	}

	middleware.JsonExit(req, 0, g.I18n().T(ctx, `{#subscribe_playlist_sucess}`), nil)
}

func (ctl *controller) AddFeedItemToPlaylist(req *ghttp.Request) {
	var (
		ctx     = req.GetCtx()
		err     error
		reqData *AddFeedItemToPlaylistReqData
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	err = feed.AddFeedItemToPlaylist(ctx, reqData.PlaylistId, reqData.ChannelId, reqData.Guid, reqData.Source)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(req.GetCtx(), `{#already_added_to_playlist}`), nil)
		} else {
			middleware.JsonExit(req, 1, err.Error(), nil)
		}
	}

	middleware.JsonExit(req, 0, g.I18n().T(ctx, `{#add_feed_item_to_playlist_sucess}`), nil)

}

func (ctl *controller) GetUserPlaylistsByUserId(req *ghttp.Request) {
	var (
		ctx             = req.GetCtx()
		err             error
		userPlaylistDto []dto.UserPlaylistDto
	)

	userId := req.Get("userId").String()
	offset := req.GetQuery("offset").Int()
	limit := req.GetQuery("limit").Int()
	userPlaylistDto, err = feed.GetUserAllPlaylists(ctx, userId, offset, limit)

	if err != nil {
		g.Log().Line().Error(ctx, err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	} else {
		middleware.JsonExit(req, 0, "", userPlaylistDto)
	}
}

func (ctl *controller) GetUserPlaylistItemList(req *ghttp.Request) {
	var (
		ctx = req.GetCtx()
		err error
	)

	userId := req.Get("userId").String()
	playlistId := req.Get("playlistId").String()
	offset := req.GetQuery("offset").Int()
	limit := req.GetQuery("limit").Int()

	itemList, err := feed.GetUserPlaylistItemList(ctx, userId, playlistId, offset, limit)

	if err != nil {
		g.Log().Line().Error(ctx, err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	} else {
		middleware.JsonExit(req, 0, "", itemList)
	}
}
