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
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
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

	excludeFeedIdList := garray.NewStrArray().Append(gstr.Split(excludeFeedId, ",")...)
	for _, searchResult := range searchResultList {
		if excludeFeedIdList.Contains(searchResult.CollectionId) {
			continue
		}
		itemID := GenerateFeedItemId(searchResult.FeedUrl, searchResult.TrackName)
		channelID := GenerateFeedChannelId(searchResult.FeedUrl, searchResult.CollectionName)
		itemList = append(itemList, dto.FeedItem{
			Id:              itemID,
			GUID:            searchResult.EpisodeGuid,
			FeedId:          searchResult.CollectionId,
			ChannelId:       channelID,
			ChannelTitle:    searchResult.CollectionName,
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

	// reverse order feedItemList
	for i, j := 0, len(feedItemList)-1; i < j; i, j = i+1, j-1 {
		feedItemList[i], feedItemList[j] = feedItemList[j], feedItemList[i]
	}

	for _, item := range feedItemList {
		model := entity.FeedItem{
			Id:              item.Id,
			ChannelId:       item.ChannelId,
			ChannelTitle:    item.ChannelTitle,
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
			Source:          item.Source,
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

func LookupItunesFeedItem(ctx context.Context, collectionId, guid string) (item dto.FeedItem, err error) {

	var (
		itunesLookupAPI = "https://itunes.apple.com/lookup?entity=podcast&id=%s"
		apiUrl          string
		decodeGUID      string
	)

	apiUrl = fmt.Sprintf(itunesLookupAPI, gurl.Encode(collectionId))
	respStr := network.GetContent(ctx, apiUrl)
	respJson := gjson.New(respStr)
	if respJson.IsNil() {
		err = errors.New("lookup result is nil")
		return
	}

	resultsJsonList := respJson.GetJsons("results")
	if respJson.IsNil() {
		err = errors.New("lookup result is nil")
		return
	}

	decodeGUID, err = gurl.Decode(guid)
	if err != nil {
		return
	}

	var feedLink string
	for index, resultsJson := range resultsJsonList {
		if index == 0 {
			var lookupResult ItunesSearchPodcastResult
			resultsJson.Scan(&lookupResult)
			feedLink = lookupResult.FeedUrl
			break
		}
	}

	item, err = GetFeedItemFromFeedLink(ctx, feedLink, decodeGUID)

	return
}

func GetFeedItemFromFeedLink(ctx context.Context, feedLink, guid string) (item dto.FeedItem, err error) {
	var (
		feed *gofeed.Feed
	)
	respStr := network.GetContent(ctx, feedLink)
	feed = ParseFeed(ctx, respStr)
	if feed != nil {
		for _, feedItem := range feed.Items {
			if feedItem.GUID == guid {
				itemID := GenerateFeedItemId(feedLink, feedItem.Title)
				channelID := GenerateFeedChannelId(feedLink, feed.Title)
				item = dto.FeedItem{
					Id:          itemID,
					ChannelId:   channelID,
					GUID:        guid,
					Title:       feedItem.Title,
					Link:        feedItem.Link,
					InputDate:   gtime.Now(),
					Duration:    feedItem.ITunesExt.Duration,
					Episode:     feedItem.ITunesExt.Episode,
					EpisodeType: feedItem.ITunesExt.EpisodeType,
					Season:      feedItem.ITunesExt.Season,
					Description: feedItem.Description,
				}

				if feedItem.PublishedParsed != nil {
					item.PubDate = feedItem.PublishedParsed.Format("Y-m-d H:i:s")
				} else {
					item.PubDate = feedItem.Published
				}

				if feedItem.Image != nil {
					item.ImageUrl = feedItem.Image.URL
				}

				if feedItem.Authors != nil || len(feedItem.Authors) > 0 {
					var (
						authors []string
					)
					for _, author := range feedItem.Authors {
						author.Name = formatFeedAuthor(author.Name)
						authors = append(authors, author.Name)
					}
					if len(authors) == 0 {
						item.Author = authors[0]
					} else {
						item.Author = gstr.Join(authors, ",")
					}
				}

				if len(feedItem.Enclosures) > 0 && feedItem.Enclosures[0] != nil {
					item.EnclosureUrl = feedItem.Enclosures[0].URL
					item.EnclosureType = feedItem.Enclosures[0].Type
					item.EnclosureLength = feedItem.Enclosures[0].Length
				}
			}
		}
	}
	return
}
