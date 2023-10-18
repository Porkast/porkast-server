package feed

import (
	"context"
	"errors"
	"fmt"
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/cache"
	"porkast-server/internal/service/elasticsearch"
	"porkast-server/internal/service/internal/dao"
	"porkast-server/internal/service/network"
	"porkast-server/utility"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mmcdole/gofeed"
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
		// feedItemInfoDto.Description = formatItemShownotes(feedItemInfoDto.Description)
	}

	if feedItemInfoDto.TextDescription == "" && feedItemInfoDto.Description != "" {
		rootDocs := soup.HTMLParse(feedItemInfoDto.Description)
		feedItemInfoDto.TextDescription = rootDocs.FullText()
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

func SearchPodcastEpisodesFromItunes(ctx context.Context, keyword, country, excludeFeedId string) (itemList []dto.FeedItem, err error) {
	var (
		ituneSearchAPI   = "https://itunes.apple.com/search?term=%s&entity=%s&media=podcast&country=%s&limit=200"
		apiUrl           string
		searchResultList []ItunesSearchEpisodeResult
	)

	if country == "" {
		country = "US"
	}

	apiUrl = fmt.Sprintf(ituneSearchAPI, gurl.Encode(keyword), "podcastEpisode", country)
	respStr := network.GetContent(ctx, apiUrl)
	respJson := gjson.New(respStr)
	if respJson.IsNil() {
		err = errors.New("search result is nil")
		return
	}

	resultsJson := respJson.Get("results")
	if respJson.IsNil() {
		err = errors.New("search result is nil")
		return
	}
	searchResultList = make([]ItunesSearchEpisodeResult, 0)
	resultsJson.Scan(&searchResultList)

	for _, searchResult := range searchResultList {
		if searchResult.CollectionId == excludeFeedId {
			continue
		}
		itemID := GenerateFeedItemId(searchResult.FeedUrl, searchResult.TrackName)
		channelID := GenerateFeedChannelId(searchResult.FeedUrl, searchResult.CollectionName)
		itemList = append(itemList, dto.FeedItem{
			Id:              itemID,
			GUID:            searchResult.EpisodeGuid,
			FeedId:          searchResult.CollectionId,
			ChannelId:       channelID,
			Source:          "itunes",
			Title:           searchResult.TrackName,
			HighlightTitle:  searchResult.TrackName,
			Link:            searchResult.TrackViewUrl,
			PubDate:         searchResult.ReleaseDate,
			ImageUrl:        searchResult.ArtworkUrl60,
			EnclosureUrl:    searchResult.EpisodeUrl,
			EnclosureType:   searchResult.EpisodeContentType,
			EnclosureLength: "",
			Duration:        gconv.String(searchResult.TrackTimeMillis),
			Description:     searchResult.Description,
			TextDescription: searchResult.Description,
			FeedLink:        searchResult.FeedUrl,
			HasThumbnail:    true,
		})
	}

	return
}

func BatchStoreFeedItems(ctx context.Context, feedItemList []dto.FeedItem) (err error) {

	for _, item := range feedItemList {
		model := entity.FeedItem{
			Id:              item.Id,
			ChannelId:       item.ChannelId,
			Guid:            item.GUID,
			Title:           item.Title,
			Link:            item.Link,
			PubDate:         gtime.New(item.PubDate),
			Author:          item.Author,
			InputDate:       gtime.New(item.InputDate),
			ImageUrl:        item.ImageUrl,
			EnclosureUrl:    item.EnclosureUrl,
			EnclosureType:   item.EnclosureType,
			EnclosureLength: item.EnclosureLength,
			Duration:        item.Duration,
			Episode:         item.Episode,
			Explicit:        item.Explicit,
			Season:          item.Season,
			EpisodeType:     item.EpisodeType,
			Description:     item.Description,
			FeedId:          item.FeedId,
			FeedLink:        item.FeedLink,
		}
		err = dao.InsertFeedItemIfNotExist(ctx, model)
		if err != nil && err.Error() == consts.DB_DATA_ALREADY_EXIST {
			err = nil
		}
	}

	return
}

func GetFeedItemsByFeedLink(ctx context.Context, feedLink string) (feed *gofeed.Feed, err error) {
	respStr := network.GetContent(ctx, feedLink)
	if respStr == "" {
		g.Log().Line().Error(ctx, "Get Feed Items By Feed Link Error")
		err = errors.New("Get Feed Items By Feed Link Error")
		return
	}

	if utility.IsStringRSSXml(respStr) {
		feed = utility.ParseFeed(ctx, respStr)
	} else {
		g.Log().Line().Error(ctx, "The Feed Is Not RSS, feed link is %s", feedLink)
		err = errors.New("The Feed Is Not RSS")
		return
	}

	return
}
