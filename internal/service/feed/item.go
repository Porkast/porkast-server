package feed

import (
	"context"
	"fmt"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/elasticsearch"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
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
	if feedItemInfoDto.ChannelImageUrl != "" {
		feedItemInfoDto.HasThumbnail = true
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
		itemDto.ChannelImageUrl = feedItemES.ChannelImageUrl
		itemDto.ChannelTitle = feedItemES.ChannelTitle
		itemDto.FeedLink = feedItemES.FeedLink
		if itemDto.ChannelImageUrl != "" {
			itemDto.HasThumbnail = true
		} else {
			itemDto.HasThumbnail = false
		}
		itemDto.PubDate = formatPubDate(itemDto.PubDate)
		itemDto.Duration = formatDuration(itemDto.Duration)
		items = append(items, itemDto)
	}

	return
}

func formatPubDate(pubDate string) (formatPubDate string) {
	formatPubDate = gtime.New(pubDate).Format("Y-m-d")
	return
}

func formatDuration(duration string) (formatDuration string) {
	if !gstr.Contains(duration, ":") {
		var (
			totalSecs = gconv.Int(duration)
			hours     int
			minutes   int
			seconds   int
		)
		hours = totalSecs / 3600
		minutes = (totalSecs % 3600) / 60
		seconds = totalSecs % 60
		formatDuration = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	} else {
		var (
			splits []string
		)
		splits = gstr.Split(duration, ":")
		if len(splits) < 3 {
			formatDuration = "00:" + duration
		}
	}
	return
}
