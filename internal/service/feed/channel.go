package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

func GetChannelInfoByChannelId(ctx context.Context, channelId string) (feedInfo dto.FeedChannel, err error) {
	var (
		feedChannelInfo entity.FeedChannel
		feedItemList    []entity.FeedItem
	)

	feedChannelInfo, err = dao.GetFeedChannelInfoByChannelId(ctx, channelId)
	if err != nil {
		return
	}
	gconv.Struct(feedChannelInfo, &feedInfo)

	feedItemList, err = dao.GetFeedItemsByChannelId(ctx, channelId)
	if err != nil {
		return
	}

	for _, item := range feedItemList {
		var (
			feedItemDto dto.FeedItem
		)
		gconv.Struct(item, &feedItemDto)
        feedItemDto.ChannelImageUrl = feedInfo.ImageUrl
        feedItemDto.ChannelTitle = feedInfo.Title
		if feedItemDto.ChannelImageUrl != "" {
			feedItemDto.HasThumbnail = true
		}
        feedInfo.Items = append(feedInfo.Items, feedItemDto)
	}

	return
}
