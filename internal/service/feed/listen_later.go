package feed

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

func CreateListenLaterByUserIdAndFeedId(ctx context.Context, userId, channelId, itemId string) (err error) {
	var (
		newEntity entity.UserListenLater
	)

	newEntity = entity.UserListenLater{
		Id:        guid.S(),
		UserId:    userId,
		ChannelId: channelId,
		ItemId:    itemId,
		RegDate:   gtime.Now(),
	}

	err = dao.CreateListenLaterByUserIdAndFeedId(ctx, newEntity)

	return
}

func GetListenLaterByUserIdAndFeedId(ctx context.Context, userId, channelId, itemId string) (userListenLaterDto dto.UserListenLater, err error) {

	var (
		userListenLaterEntity entity.UserListenLater
		feedItemInfoEntity    entity.FeedItem
		feedItemInfoDto       dto.FeedItem
	)

	if userId == "" || channelId == "" || itemId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	userListenLaterEntity, err = dao.GetListenLaterByUserIdAndFeedId(ctx, userId, channelId, itemId)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	}
	gconv.Struct(userListenLaterEntity, &userListenLaterDto)

	feedItemInfoEntity, err = dao.GetFeedItemById(ctx, channelId, itemId)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	}
	gconv.Struct(feedItemInfoEntity, &feedItemInfoDto)

	return
}

func GetListenLaterListByUserId(ctx context.Context, userId string, offset, limit int) (userListenLaterDtoList []dto.UserListenLater, err error) {
	var (
		totalCount int
	)

	userListenLaterDtoList, err = dao.GetListenLaterListByUserId(ctx, userId, offset, limit)
	if err != nil {
		return
	}

	totalCount, err = dao.GetTotalListenLaterCountByUserId(ctx, userId)
	if err != nil {
		return
	}

	for i := 0; i < len(userListenLaterDtoList); i++ {
		var dtoItem = userListenLaterDtoList[i]
		if dtoItem.ImageUrl == "" {
			if dtoItem.ChannelImageUrl == "" {
				dtoItem.HasThumbnail = false
			} else {
				dtoItem.ImageUrl = dtoItem.ChannelImageUrl
				dtoItem.HasThumbnail = true
			}
		} else {
			dtoItem.HasThumbnail = true
		}

		if dtoItem.TextDescription == "" && dtoItem.Description != "" {
			rootDocs := soup.HTMLParse(dtoItem.Description)
			dtoItem.TextDescription = rootDocs.FullText()
		}
		if dtoItem.Author == "" {
			dtoItem.Author = dtoItem.ChannelAuthor
		}
		dtoItem.PubDate = formatPubDate(dtoItem.PubDate)
		dtoItem.Duration = formatDuration(dtoItem.Duration)
		dtoItem.RegDate = consts.ADD_ON_TEXT + formatPubDate(dtoItem.RegDate)
		dtoItem.Count = totalCount
        userListenLaterDtoList[i] = dtoItem
	}

	return
}
