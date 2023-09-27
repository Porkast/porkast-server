// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"errors"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao/internal"

	"github.com/gogf/gf/v2/frame/g"
)

// feedItemDao is the data access object for table feed_item.
// You can define custom methods on it to extend its functionality as you wish.
type feedItemDao struct {
	*internal.FeedItemDao
}

var (
	// FeedItem is globally public accessible object for table feed_item operations.
	FeedItem = feedItemDao{
		internal.NewFeedItemDao(),
	}
)

// Fill with you ideas below.
func GetFeedItemsByChannelId(ctx context.Context, channelId string, offset, limit int) (itemList []entity.FeedItem, err error) {

	if limit == 0 {
		limit = 10
	}
	err = FeedItem.Ctx(ctx).Where("channel_id=?", channelId).Offset(offset).Limit(limit).OrderDesc("pub_date").Scan(&itemList)
	if err != nil {
		return
	}

	if len(itemList) == 0 {
		return itemList, errors.New("The feed item is exist.")
	}

	return
}

func GetFeedItemById(ctx context.Context, channelId, itemId string) (item entity.FeedItem, err error) {

	err = FeedItem.Ctx(ctx).Where("id=?", itemId).Where("channel_id=?", channelId).Scan(&item)
	if err != nil {
		return
	}

	return
}

func GetFeedItemCountByChannelId(ctx context.Context, channelId string) (count int, err error) {

	count, err = FeedItem.Ctx(ctx).Where("channel_id=?", channelId).Count()

	return
}

func GetLatestPubFeedItems(ctx context.Context, offset, limit int) (entities []dto.FeedItem) {
	if limit == 0 {
		limit = 10
	}

	g.Model("feed_item fi").
		InnerJoin("feed_channel fc", "fc.id=fi.channel_id").
		Fields("fi.*, fc.image_url as channel_image_url, fc.feed_link, fc.title as channel_title, fc.author as channelAuthor").
		Offset(offset).
		Limit(limit).
		OrderDesc("fi.pub_date").
		Scan(&entities)

	return
}

func GetFeedItemListByPubDate(ctx context.Context, startDate, endDate string) (entities []dto.FeedItem) {

	g.Model("feed_item fi").
		InnerJoin("feed_channel fc", "fc.id=fi.channel_id").
		Fields("fi.*, fc.image_url as channel_image_url, fc.feed_link, fc.title as channel_title, fc.author as channelAuthor").
		Where("fi.pub_date>=?", startDate).
		Where("fi.pub_date<?", endDate).
		Scan(&entities)

	return
}

func GetFeedChannelItemListByPubDate(ctx context.Context, channelId, pubDate string) (entities []dto.FeedItem) {

	g.Model("feed_item fi").
		InnerJoin("feed_channel fc", "fc.id=fi.channel_id").
		Fields("fi.*, fc.image_url as channel_image_url, fc.feed_link, fc.title as channel_title, fc.author as channelAuthor").
		Where("fi.pub_date=? and fc.id=?", pubDate, channelId).
		Scan(&entities)

	return
}
