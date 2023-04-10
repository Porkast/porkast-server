// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"errors"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao/internal"
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
func GetFeedItemsByChannelId(ctx context.Context, channelId string) (itemList []entity.FeedItem, err error) {

	err = FeedItem.Ctx(ctx).Where("channel_id=?", channelId).Scan(&itemList)
	if err != nil {
		return
	}

	if len(itemList) == 0 {
		return itemList, errors.New("The feed item is exist.")
	}

	return
}

func GetFeedItemById(ctx context.Context, id string) (item entity.FeedItem, err error) {

	err = FeedItem.Ctx(ctx).Where("id=?", id).Scan(&item)
	if err != nil {
		return
	}

	return
}