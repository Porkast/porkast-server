package feed

import (
	"context"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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

	err = dao.InsertNewUserPlaylistIfNotExist(ctx, entity)

	return
}

func SubscribePlaylist(ctx context.Context, userId, playlistId string) (err error) {

	if userId == "" || playlistId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	existEntity, err := dao.GetPlaylistById(ctx, playlistId)
	if err != nil {
		return gerror.New("Playlist " + gcode.CodeNotFound.Message())
	}

	newPlaylistId := GeneratePlaylistId(existEntity.PlaylistName, userId)
	newEntity := entity.UserPlaylist{
		Id:             newPlaylistId,
		OrigPlaylistId: existEntity.UserId,
		PlaylistName:   existEntity.PlaylistName,
		Description:    existEntity.Description,
		UserId:         userId,
		CreatorId:      existEntity.UserId,
		RegDate:        gtime.Now(),
		Status:         1,
	}

	err = dao.InsertNewUserPlaylistIfNotExist(ctx, newEntity)

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
		RegDate:    gtime.Now(),
		Status:     1,
	}

	err = dao.InsertNewUserPlaylistItemIfNotExit(ctx, entity)

	return
}

func GetUserAllPlaylists(ctx context.Context, userId string, offset, limit int) (result []dto.UserPlaylistDto, err error) {

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	var (
		entities   []entity.UserPlaylist
		totalCount int
	)

	if limit == 0 {
		limit = 10
	}

	entities, err = dao.GetUserPlaylistsByUserId(ctx, userId, offset, limit)
	gconv.Structs(entities, &result)

	totalCount, err = dao.GetUserPlaylistTotalCountByUserId(ctx, userId)
	for index := range result {
		result[index].Count = totalCount
	}

	return
}

func GetUserPlaylistItemList(ctx context.Context, userId, playlistId string, offset, limit int) (result []dto.UserPlaylistItemDto, err error) {

	if userId == "" || playlistId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	if limit == 0 {
		limit = 10
	}

	result, err = dao.GetUserPlaylistItemsById(ctx, playlistId, offset, limit)
	if err != nil {
		return nil, err
	}

	totalCount, err := dao.GetUserPlaylistItemTotalCount(ctx, playlistId)
	if err != nil {
		return nil, err
	}

	for index := range result {
		result[index].Count = totalCount
		result[index].Duration = formatDuration(result[index].Duration)
	}

	return
}
