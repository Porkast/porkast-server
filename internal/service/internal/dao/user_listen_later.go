// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao/internal"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// userListenLaterDao is the data access object for table user_listen_later.
// You can define custom methods on it to extend its functionality as you wish.
type userListenLaterDao struct {
	*internal.UserListenLaterDao
}

var (
	// UserListenLater is globally public accessible object for table user_listen_later operations.
	UserListenLater = userListenLaterDao{
		internal.NewUserListenLaterDao(),
	}
)

// Fill with you ideas below.
func CreateListenLaterByUserIdAndFeedId(ctx context.Context, newEntity entity.UserListenLater) (err error) {
	var (
		queryEntity entity.UserListenLater
	)

	queryEntity, err = GetListenLaterByUserIdAndFeedId(ctx, newEntity.UserId, newEntity.ChannelId, newEntity.ItemId)
	if queryEntity.Id != "" {
		err = gerror.New(consts.DB_DATA_ALREADY_EXIST)
		return
	}

	if newEntity.UserId == "" || newEntity.ChannelId == "" || newEntity.ItemId == "" || newEntity.RegDate == nil || newEntity.RegDate.IsDST() {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	_, err = UserListenLater.Ctx(ctx).Insert(newEntity)

	return
}

func GetListenLaterByUserIdAndFeedId(ctx context.Context, userId, channelId, itemId string) (entity entity.UserListenLater, err error) {

	if userId == "" || channelId == "" || itemId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}
	err = UserListenLater.Ctx(ctx).Where("user_id=? and channel_id=? and item_id=? and status=1", userId, channelId, itemId).Scan(&entity)

	return
}

func GetListenLaterListByUserId(ctx context.Context, userId string, offset, limit int) (dtoList []dto.UserListenLater, err error) {

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	g.Model("user_listen_later ull").
		InnerJoin("feed_item fi", "ull.item_id = fi.id").
		InnerJoin("feed_channel fc", "fi.channel_id = fc.id").
		Fields("ull.*, fi.title, fi.link, fi.pub_date, fi.author, fi.input_date, fi.image_url, fi.enclosure_url, fi.enclosure_type, fi.enclosure_length, fi.duration, fi.episode, fi.explicit, fi.season, fi.episodeType, fi.description, fc.image_url as channel_image_url, fc.feed_link, fc.title as channel_title, fc.author as channelAuthor").
		Where("status=1").
		Offset(offset).
		Limit(limit).
		Order("ull.reg_date desc").
		Scan(&dtoList)

	return
}

func GetTotalListenLaterCountByUserId(ctx context.Context, userId string) (totalCount int, err error) {

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	totalCount, err = UserListenLater.Ctx(ctx).Where("user_id=? and status=1", userId).Count()

	return
}
