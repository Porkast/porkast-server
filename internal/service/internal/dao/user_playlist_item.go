// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao/internal"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// userPlaylistItemDao is the data access object for table user_playlist_item.
// You can define custom methods on it to extend its functionality as you wish.
type userPlaylistItemDao struct {
	*internal.UserPlaylistItemDao
}

var (
	// UserPlaylistItem is globally public accessible object for table user_playlist_item operations.
	UserPlaylistItem = userPlaylistItemDao{
		internal.NewUserPlaylistItemDao(),
	}
)

// Fill with you ideas below.
func InsertNewUserPlaylistItemIfNotExit(ctx context.Context, newEntity entity.UserPlaylistItem) (err error) {

	var (
		result gdb.Record
	)

	result, err = UserPlaylistItem.Ctx(ctx).Where("id=?", newEntity.Id).One()
	if err != nil {
		return
	}

	if !result.IsEmpty() {
		err = gerror.New(consts.DB_DATA_ALREADY_EXIST)
		return
	}

	_, err = UserPlaylistItem.Ctx(ctx).Insert(newEntity)
	return
}

func GetUserPlaylistItemsById(ctx context.Context, playlistId string, offset, limit int) (entities []dto.UserPlaylistItemDto, err error) {
	g.Model("user_playlist_item upi").
		LeftJoin("feed_item fi", "fi.id=upi.item_id").
		Where("upi.playlist_id=? and upi.status=1", playlistId).
		Fields("fi.*, upi.reg_date").
		OrderDesc("upi.reg_date").
		Offset(offset).
		Limit(limit).
		Scan(&entities)
	return
}

func GetUserPlaylistItemTotalCount(ctx context.Context, playlistId string) (totalCount int, err error) {

	if playlistId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	totalCount, err = UserPlaylistItem.Ctx(ctx).Where("playlist_id=? and status=1", playlistId).Count()

	return
}
