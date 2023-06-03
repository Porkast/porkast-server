package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/elasticsearch"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetChannelInfoByChannelId(ctx context.Context, channelId string, offset, limit int) (feedInfo dto.FeedChannel, err error) {
	g.Log().Line().Debugf(ctx, "GetChannelInfoByChannelId channelId : %s, offset : %d, limit: %d", channelId, offset, limit)
	var (
		feedChannelInfo entity.FeedChannel
		feedItemList    []entity.FeedItem
		totalCount      int
	)

	feedChannelInfo, err = dao.GetFeedChannelInfoByChannelId(ctx, channelId)
	if err != nil {
		return
	}
	gconv.Struct(feedChannelInfo, &feedInfo)

	feedItemList, err = dao.GetFeedItemsByChannelId(ctx, channelId, offset, limit)
	if err != nil {
		return
	}

	totalCount, err = dao.GetFeedItemCountByChannelId(ctx, channelId)
	if err != nil {
		return
	}
	feedInfo.Count = totalCount

	for _, item := range feedItemList {
		var (
			feedItemDto dto.FeedItem
		)
		gconv.Struct(item, &feedItemDto)
		feedItemDto.ChannelImageUrl = feedInfo.ImageUrl
		feedItemDto.ChannelTitle = feedInfo.Title
		feedItemDto.Duration = formatDuration(feedItemDto.Duration)
		feedItemDto.PubDate = formatPubDate(feedItemDto.PubDate)
		if feedItemDto.ImageUrl != "" {
			feedItemDto.HasThumbnail = true
		} else if feedItemDto.ChannelImageUrl != "" {
			feedItemDto.ImageUrl = feedItemDto.ChannelImageUrl
			feedItemDto.HasThumbnail = true
		} else {
			feedItemDto.HasThumbnail = false
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

func QueryFeedChannelByKeyword(ctx context.Context, keyword string, page, size int) (esChannelList []dto.FeedChannel, err error) {
	var (
		esChannelEntityList []entity.FeedChannelESData
	)

	if size == 0 {
		size = 10
	}

	if page >= 1 {
		page = (page - 1) * size
	} else {
		page = page * size
	}

	esChannelEntityList, err = elasticsearch.GetClient().QueryFeedChannelFull(ctx, keyword, page, size)
	if err != nil {
		return
	}

	for _, esChannelEntity := range esChannelEntityList {
		var esChannelDto dto.FeedChannel
		gconv.Struct(esChannelEntity, &esChannelDto)
		if esChannelDto.ImageUrl == "" {
			esChannelDto.HasThumbnail = false
		} else {
			esChannelDto.HasThumbnail = true
		}
		esChannelList = append(esChannelList, esChannelDto)
	}

	return
}
