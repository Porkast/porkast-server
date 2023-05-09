package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/anaskhan96/soup"
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
		feedItemDto.Duration = formatDuration(feedItemDto.Duration)
		feedItemDto.PubDate = formatPubDate(feedItemDto.PubDate)
		if feedItemDto.ChannelImageUrl != "" {
			feedItemDto.HasThumbnail = true
		}
        if feedItemDto.HighlightTitle == "" {
            feedItemDto.HighlightTitle = feedItemDto.Title
        }
        if feedItemDto.TextDescription == "" && feedItemDto.Description != "" {
			rootDocs := soup.HTMLParse(feedItemDto.Description)
			feedItemDto.TextDescription = rootDocs.FullText()
        }
		feedInfo.Items = append(feedInfo.Items, feedItemDto)
	}

	return
}
