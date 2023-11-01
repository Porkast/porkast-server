package ctls

import (
	"porkast-server/internal/consts"
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
		middleware.JsonExit(req, 1, err.Error(), nil)
	}

	middleware.JsonExit(req, 0, g.I18n().T(ctx, `{#create_playlist_sucess}`), nil)

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
