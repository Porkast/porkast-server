package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/elasticsearch"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
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

func SearchFeedItemsByKeyword(ctx context.Context, keyword string, page, size int) (items []dto.FeedItem, err error) {
	var (
		feedItemESDatalList []entity.FeedItemESData
	)

	if size == 0 {
		size = 10
	}

	if page >= 1 {
		page = (page - 1) * size
	} else {
		page = page * size
	}

	feedItemESDatalList, err = elasticsearch.GetClient().QueryFeedItemFull(ctx, keyword, page, size)
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
		items = append(items, itemDto)
	}

	g.Log().Line().Debug(ctx, "search result :\n", gjson.MustEncodeString(feedItemESDatalList))

	return
}
