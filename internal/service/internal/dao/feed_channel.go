// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao/internal"
)

// feedChannelDao is the data access object for table feed_channel.
// You can define custom methods on it to extend its functionality as you wish.
type feedChannelDao struct {
	*internal.FeedChannelDao
}

var (
	// FeedChannel is globally public accessible object for table feed_channel operations.
	FeedChannel = feedChannelDao{
		internal.NewFeedChannelDao(),
	}
)

// Fill with you ideas below.

func GetFeedChannelInfoByChannelId(ctx context.Context, channelId string) (feedInfo entity.FeedChannel, err error) {

	err = FeedChannel.Ctx(ctx).Where("id=?", channelId).Scan(&feedInfo)

	return
}

func GetZHFeedChannelTotalCount(ctx context.Context) (count int, err error) {
	count, err = FeedChannel.Ctx(ctx).Where("language like 'zh%'").Count()
	return
}

func GetZHFeedChannelList(ctx context.Context) (entities []entity.FeedChannel, err error) {

	err = FeedChannel.Ctx(ctx).Where("language like 'zh%'").Scan(&entities)

	return
}
