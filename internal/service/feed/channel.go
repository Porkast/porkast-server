package feed

import (
	"context"
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/cache"
	"porkast-server/internal/service/elasticsearch"
	"porkast-server/internal/service/internal/dao"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
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
	feedInfo.Author = formatFeedAuthor(feedInfo.Author)
	if feedInfo.ImageUrl == "" {
		feedInfo.HasThumbnail = false
	} else {
		feedInfo.HasThumbnail = true
	}

	if feedChannelInfo.Categories != "" {
		feedInfo.Categories = gstr.Split(feedChannelInfo.Categories, ",")
	}

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
		feedItemDto.Author = formatFeedAuthor(feedItemDto.Author)
        feedItemDto.Title = formatItemTitle(feedItemDto.Title)
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

func QueryFeedChannelByKeyword(ctx context.Context, params SearchParams) (esChannelList []dto.FeedChannel, err error) {
	var (
		esChannelEntityList []entity.FeedChannelESData
	)

	if params.Size == 0 {
		params.Size = 10
	}

	if params.Page >= 1 {
		params.Page = (params.Page - 1) * params.Size
	} else {
		params.Page = params.Page * params.Size
	}

	esChannelEntityList, err = elasticsearch.GetClient().QueryFeedChannelFull(ctx, params.Keyword, params.Page, params.Size)
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
		if esChannelEntity.Categories != "" {
			esChannelDto.Categories = gstr.Split(esChannelEntity.Categories, ",")
		}
		esChannelDto.Author = formatFeedAuthor(esChannelDto.Author)
		esChannelList = append(esChannelList, esChannelDto)
	}

	return
}

func GetAllFeedChannelCount(ctx context.Context) (count int, err error) {
	var (
		countVar *gvar.Var
	)

	countVar, err = cache.GetCache(ctx, gconv.String(consts.FEED_CHANNEL_TOTAL_COUNT))
	if err == nil && countVar != nil && countVar.Int() != 0 {
		count = countVar.Int()
		return
	} else {
		count, err = dao.GetZHFeedChannelTotalCount(ctx)
	}

	return
}

func GetAllFeedChannelCountFromCache(ctx context.Context) (count int, err error) {
	var (
		countVar *gvar.Var
	)

	countVar, err = cache.GetCache(ctx, gconv.String(consts.FEED_CHANNEL_TOTAL_COUNT))
	if countVar != nil {
		count = countVar.Int()
	}

	return
}
