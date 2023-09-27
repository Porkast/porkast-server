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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetFeedItemByItemId(ctx context.Context, channelId, itemId string) (feedChannelDto dto.FeedChannel, feedItemInfoDto dto.FeedItem, err error) {
	var (
		feedItemModel    entity.FeedItem
		feedChannelModel entity.FeedChannel
	)

	feedChannelModel, err = dao.GetFeedChannelInfoByChannelId(ctx, channelId)
	if err != nil {
		return
	}
	gconv.Struct(feedChannelModel, &feedChannelDto)
	feedItemModel, err = dao.GetFeedItemById(ctx, channelId, itemId)
	gconv.Struct(feedItemModel, &feedItemInfoDto)
	feedItemInfoDto.Duration = formatDuration(feedItemInfoDto.Duration)
	feedItemInfoDto.PubDate = formatPubDate(feedItemInfoDto.PubDate)
	feedItemInfoDto.ChannelImageUrl = feedChannelModel.ImageUrl
	feedItemInfoDto.ChannelTitle = feedChannelModel.Title
	feedItemInfoDto.FeedLink = feedChannelModel.FeedLink
	feedItemInfoDto.Link = formatSourceLink(feedItemInfoDto.Link)
	feedItemInfoDto.Title = formatItemTitle(feedItemInfoDto.Title)
	if feedItemInfoDto.ImageUrl == "" {
		if feedItemInfoDto.ChannelImageUrl != "" {
			feedItemInfoDto.ImageUrl = feedItemInfoDto.ChannelImageUrl
			feedItemInfoDto.HasThumbnail = true
		} else {
			feedItemInfoDto.HasThumbnail = false
		}
	} else {
		feedItemInfoDto.HasThumbnail = true
	}

	if feedItemInfoDto.Author == "" {
		feedItemInfoDto.Author = feedChannelDto.Author
	}

	feedItemInfoDto.Author = formatFeedAuthor(feedItemInfoDto.Author)

	if feedItemInfoDto.Description != "" {
		feedItemInfoDto.Description = formatItemShownotes(feedItemInfoDto.Description)
	}

	return
}

func SearchFeedItemsByKeyword(ctx context.Context, params SearchParams) (items []dto.FeedItem, err error) {
	var (
		feedItemESDatalList []entity.FeedItemESData
	)

	if params.Size == 0 {
		params.Size = 10
	}

	if params.Page >= 1 {
		params.Page = (params.Page - 1) * params.Size
	} else {
		params.Page = params.Page * params.Size
	}

	feedItemESDatalList, err = elasticsearch.GetClient().QueryFeedItemFull(ctx, params.Keyword, params.SortByDate, params.Page, params.Size)
	if err != nil {
		return
	}
	for _, feedItemES := range feedItemESDatalList {
		var itemDto dto.FeedItem
		gconv.Struct(feedItemES, &itemDto)
		itemDto.Link = formatSourceLink(itemDto.Link)
		if itemDto.ImageUrl != "" {
			itemDto.HasThumbnail = true
		} else if itemDto.ChannelImageUrl != "" {
			itemDto.ImageUrl = itemDto.ChannelImageUrl
			itemDto.HasThumbnail = true
		} else {
			itemDto.HasThumbnail = false
		}
		if itemDto.HighlightTitle == "" {
			itemDto.HighlightTitle = itemDto.Title
		}
		if itemDto.TextDescription == "" && itemDto.Description != "" {
			rootDocs := soup.HTMLParse(itemDto.Description)
			itemDto.TextDescription = rootDocs.FullText()
		}
		itemDto.Author = formatFeedAuthor(itemDto.Author)
		itemDto.PubDate = formatPubDate(itemDto.PubDate)
		itemDto.Duration = formatDuration(itemDto.Duration)
		itemDto.HighlightChannelTitle = itemDto.ChannelTitle
		itemDto.ChannelTitle = formatTitle(itemDto.HighlightChannelTitle)
		itemDto.Title = formatItemTitle(itemDto.Title)
		items = append(items, itemDto)
	}

	// g.Log().Line().Debug(ctx, "search result :\n", gjson.MustEncodeString(feedItemESDatalList))

	return
}

func GetFeedItemCountByChannelId(ctx context.Context, channelId string) (count int, err error) {

	count, err = dao.GetFeedItemCountByChannelId(ctx, channelId)

	return
}

func GetAllFeedItemCountFromCache(ctx context.Context) (count int, err error) {
	var (
		countVar *gvar.Var
	)

	countVar, err = cache.GetCache(ctx, gconv.String(consts.FEED_ITEM_TOTAL_COUNT))
	if countVar != nil {
		count = countVar.Int()
	}
	return
}

func GetLatestPubFeedItems(ctx context.Context, offset, limit int) (itemList []dto.FeedItem, err error) {

	itemList = dao.GetLatestPubFeedItems(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(itemList); i++ {
		itemDto := itemList[i]
		itemDto.Author = formatFeedAuthor(itemDto.Author)
		itemDto.PubDate = formatPubDate(itemDto.PubDate)
		itemDto.Duration = formatDuration(itemDto.Duration)
		itemDto.Link = formatSourceLink(itemDto.Link)
		itemList[i] = itemDto
	}

	return
}

func GetPubFeedItemsByDate(ctx context.Context, date string) (itemList []dto.FeedItem, err error) {
	var (
		startDate    *gtime.Time
		startDateStr string
		endDate      *gtime.Time
		endDateStr   string
	)

	startDate = gtime.NewFromStr(date)
	endDate = gtime.NewFromStr(date).EndOfDay()

	startDateStr = startDate.ISO8601()
	endDateStr = endDate.ISO8601()

	itemList = dao.GetFeedItemListByPubDate(ctx, startDateStr, endDateStr)

	return
}
