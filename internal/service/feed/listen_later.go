package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"

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

func GetListenLaterListByUserId(ctx context.Context, userId string) (userListenLaterDtoList []dto.UserListenLater, err error) {
	var (
		userListenLaterEntityList []entity.UserListenLater
	)

	userListenLaterEntityList, err = dao.GetListenLaterListByUserId(ctx, userId)
	if err != nil {
		return
	}

	gconv.Structs(userListenLaterEntityList, &userListenLaterDtoList)
	
	return
}