package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/elasticsearch"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

func GetFeedItemByItemId(ctx context.Context, itemId string) (feedItemInfoDto dto.FeedItem, err error) {
	var (
		feedItemModel entity.FeedItem
	)

	feedItemModel, err = dao.GetFeedItemById(ctx, itemId)
	gconv.Struct(feedItemModel, &feedItemModel)

	return
}

func SearchFeedItemsByKeyword(ctx context.Context, keyword string, start, size int) (items []dto.FeedItem, err error) {
	var (
		feedItemESDatalList []entity.FeedItemESData
	)

	feedItemESDatalList, err = elasticsearch.GetClient().QueryFeedItemFull(ctx, keyword, start, size)
	if err != nil {
		return
	}
	for _, feedItemES := range feedItemESDatalList {
		var itemDto dto.FeedItem
		gconv.Struct(feedItemES, &itemDto)
		itemDto.ChannelImageUrl = feedItemES.ChannelImageUrl
		itemDto.ChannelTitle = feedItemES.Title
		itemDto.SourceLink = feedItemES.SourceLink
		if itemDto.ImageUrl != "" {
			itemDto.HasThumbnail = true
		} else {
			itemDto.HasThumbnail = false
		}
		items = append(items, itemDto)
	}

	return
}
