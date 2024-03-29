// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"porkast-server/internal/consts"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao/internal"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

// userPlaylistDao is the data access object for table user_playlist.
// You can define custom methods on it to extend its functionality as you wish.
type userPlaylistDao struct {
	*internal.UserPlaylistDao
}

var (
	// UserPlaylist is globally public accessible object for table user_playlist operations.
	UserPlaylist = userPlaylistDao{
		internal.NewUserPlaylistDao(),
	}
)

// Fill with you ideas below.

func InsertNewUserPlaylistIfNotExist(ctx context.Context, newEntity entity.UserPlaylist) (err error) {
	var (
		result gdb.Record
	)

	result, err = UserPlaylist.Ctx(ctx).Where("id=?", newEntity.Id).One()
	if err != nil {
		return
	}

	if !result.IsEmpty() {
		err = gerror.New(consts.DB_DATA_ALREADY_EXIST)
		return
	}
	_, err = UserPlaylist.Ctx(ctx).Insert(newEntity)
	return
}

func GetPlaylistById(ctx context.Context, id string) (entity entity.UserPlaylist, err error) {
	err = UserPlaylist.Ctx(ctx).Where("id=?", id).Scan(&entity)
	return
}

func GetUserPlaylistTotalCountByUserId(ctx context.Context, userId string) (count int, err error) {
	count, err = UserPlaylist.Ctx(ctx).Where("user_id=? and status=1", userId).Count()
	return
}

func GetUserPlaylistsByUserId(ctx context.Context, userId string, offset, limit int) (entities []entity.UserPlaylist, err error) {
	err = UserPlaylist.Ctx(ctx).Where("user_id=? and status=1", userId).
		Offset(offset).
		Limit(limit).
		Scan(&entities)
	return
}
