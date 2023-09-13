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

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// userSubKeywordDao is the data access object for table user_sub_keyword.
// You can define custom methods on it to extend its functionality as you wish.
type userSubKeywordDao struct {
	*internal.UserSubKeywordDao
}

var (
	// UserSubKeyword is globally public accessible object for table user_sub_keyword operations.
	UserSubKeyword = userSubKeywordDao{
		internal.NewUserSubKeywordDao(),
	}
)

// Fill with you ideas below.
func GetUserSubKeywordByUserIdAndKeyword(ctx context.Context, userId, keyword string) (resultEntity entity.UserSubKeyword, err error) {

	err = UserSubKeyword.Ctx(ctx).Where("user_id=? and keyword=? and status=1", userId, keyword).Scan(&resultEntity)

	return
}

func CreateUserSubKeyword(ctx context.Context, newEntity entity.UserSubKeyword) (err error) {
	var (
		queryEntity entity.UserSubKeyword
	)

	queryEntity, err = GetUserSubKeywordByUserIdAndKeyword(ctx, newEntity.UserId, newEntity.Keyword)
	if queryEntity.Id != "" {
		err = gerror.New(consts.DB_DATA_ALREADY_EXIST)
        return
	}

	if newEntity.UserId == "" || newEntity.Keyword == "" || newEntity.CreateTime == nil || newEntity.CreateTime.IsDST() {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	_, err = UserSubKeyword.Ctx(ctx).Insert(newEntity)

	return
}

func DoSubKeywordByUserIdAndKeyword(ctx context.Context, newUSKEntity entity.UserSubKeyword, newKSEntityList []entity.KeywordSubscription) (err error) {

	var (
		queryEntity entity.UserSubKeyword
	)

	queryEntity, err = GetUserSubKeywordByUserIdAndKeyword(ctx, newUSKEntity.UserId, newUSKEntity.Keyword)

	if queryEntity.Id != "" {
		err = gerror.New(consts.DB_DATA_ALREADY_EXIST)
		return
	}

	if newUSKEntity.UserId == "" || newUSKEntity.Keyword == "" || newUSKEntity.CreateTime == nil || newUSKEntity.CreateTime.IsDST() {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var (
			err error
		)
		_, err = tx.Save("user_sub_keyword", newUSKEntity)
		if err != nil {
			tx.Rollback()
		}

		_, err = tx.Save("keyword_subscription", newKSEntityList)
		if err != nil {
			tx.Rollback()
		}

		return err
	})

}

func GetUserSubKeywordListByUserId(ctx context.Context, userId string) (dtos []dto.UserSubKeyword, err error) {

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	UserSubKeyword.Ctx(ctx).Where("user_id=? and status=1", userId).Scan(&dtos)

	return
}

func GetUserSubKeywordListByUserIdAndKeyword(ctx context.Context, userId, keyword string) (dtos []dto.UserSubKeyword, err error) {

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	g.Model("user_sub_keyword usk").
		InnerJoin("keyword_subscription ks", "usk.keyword = ks.keyword and ks.lang = usk.lang and ks.order_by_date = usk.order_by_date").
		InnerJoin("feed_channel fc", "ks.feed_channel_id = fc.id").
		InnerJoin("feed_item fi", "ks.feed_channel_id = fi.channel_id and ks.feed_item_id = fi.id").
		Fields("usk.*, fi.id as item_id, fi.channel_id ,fi.title, fi.link, fi.pub_date, fi.author, fi.input_date, fi.image_url, fi.enclosure_url, fi.enclosure_type, fi.enclosure_length, fi.duration, fi.episode, fi.explicit, fi.season, fi.episodeType, fi.description, fc.image_url as channel_image_url, fc.feed_link, fc.title as channel_title, fc.author as channelAuthor").
		Where("usk.user_id=? and usk.keyword=? and status=1", userId, keyword).
		Scan(&dtos)

	return
}

func GetAllKindSubKeywordList(ctx context.Context, offset, limit int) (entities []entity.UserSubKeyword, err error) {

	var (
		dbModel *gdb.Model
	)

	dbModel = UserSubKeyword.Ctx(ctx).
        Fields("keyword","lang","order_by_date").
		Group("keyword", "lang", "order_by_date")

	if limit == 0 {
		err = dbModel.Scan(&entities)
	} else {
		err = dbModel.Offset(limit).Limit(limit).Scan(&entities)
	}

	return
}

func GetUserSubscriptionCount(ctx context.Context, userId string) (count int, err error) {

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	count, err = UserSubKeyword.Ctx(ctx).Where("user_id=? and status=1", userId).Count()

	return
}
