package feed

import (
	"context"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func CreatePlaylist(ctx context.Context, playlistName, userId, description string) (err error) {

	if playlistName == "" || userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	playlistId := GeneratePlaylistId(playlistName, userId)
	descBytes := []byte(description)
	entity := entity.UserPlaylist{
		Id:           playlistId,
		PlaylistName: playlistName,
		Description:  descBytes,
		UserId:       userId,
		RegDate:      gtime.Now(),
		Status:       1,
	}

	err = dao.InsertNewUserPlaylist(ctx, entity)

	return
}

func AddFeedItemToPlaylist(ctx context.Context, playlistId, channelId, guid, source string) (err error) {

	if playlistId == "" || channelId == "" || guid == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	_, err = dao.GetPlaylistById(ctx, playlistId)
	if err != nil {
		return gerror.New("Playlist " + gcode.CodeNotFound.Message())
	}

	var feedItem dto.FeedItem
	var feedItems []dto.FeedItem
	if source == "" || source == "itunes" {
		feedItem, err = LookupItunesFeedItem(ctx, channelId, guid)
		if err != nil {
			return
		}

		if feedItem.Id == "" {
			return gerror.New(gcode.CodeUnknown.Message())
		}
	}

	feedItems = make([]dto.FeedItem, 0)
	feedItems = append(feedItems, feedItem)
	err = BatchStoreFeedItems(ctx, feedItems)
	if err != nil {
		return
	}

	id := GeneratePlaylistItemId(playlistId, feedItem.Id)
	entity := entity.UserPlaylistItem{
		Id:         id,
		PlaylistId: playlistId,
		ItemId:     feedItem.Id,
		ChannelId:  channelId,
		RegDate:    gtime.New(),
		Status:     0,
	}

	err = dao.InsertNewUserPlaylistItemIfNotExit(ctx, entity)

	return
}
