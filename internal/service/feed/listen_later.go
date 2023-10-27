package feed

import (
	"context"
	"fmt"
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao"
	"porkast-server/internal/service/user"

	"github.com/anaskhan96/soup"
	"github.com/eduncan911/podcast"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

func CreateListenLaterByUserIdAndFeedId(ctx context.Context, userId, channelId, itemId, source string) (err error) {
	var (
		newEntity entity.UserListenLater
		feedItem dto.FeedItem
		feedItemEntity entity.FeedItem
	)

	if source == "" || source == "itunes" {
		feedItem, err = LookupItunesFeedItem(ctx, channelId, itemId)
	}

	feedItemEntity = entity.FeedItem{
			Id:              feedItem.Id,
			ChannelId:       feedItem.ChannelId,
			ChannelTitle:    feedItem.ChannelTitle,
			Guid:            feedItem.GUID,
			Title:           feedItem.Title,
			Link:            feedItem.Link,
			PubDate:         gtime.New(feedItem.PubDate),
			Author:          feedItem.Author,
			InputDate:       gtime.New(feedItem.InputDate),
			ImageUrl:        feedItem.ImageUrl,
			EnclosureUrl:    feedItem.EnclosureUrl,
			EnclosureType:   feedItem.EnclosureType,
			EnclosureLength: feedItem.EnclosureLength,
			Duration:        feedItem.Duration,
			Episode:         feedItem.Episode,
			Explicit:        feedItem.Explicit,
			Season:          feedItem.Season,
			EpisodeType:     feedItem.EpisodeType,
			Description:     feedItem.Description,
			FeedId:          feedItem.FeedId,
			FeedLink:        feedItem.FeedLink,
			Source:          feedItem.Source,
	}

	err = dao.InsertFeedItemIfNotExist(ctx, feedItemEntity)
	if err != nil && err.Error() != consts.DB_DATA_ALREADY_EXIST {
		return
	}

	newEntity = entity.UserListenLater{
		Id:        guid.S(),
		UserId:    userId,
		ChannelId: feedItem.ChannelId,
		ItemId:    feedItem.Id,
		Status:    1,
		RegDate:   gtime.Now(),
	}

	err = dao.CreateListenLaterByUserIdAndFeedId(ctx, newEntity)

	return
}

func GetListenLaterByUserIdAndFeedId(ctx context.Context, userId, channelId, itemId string) (userListenLaterDto dto.UserListenLater, err error) {

	var (
		userListenLaterEntity entity.UserListenLater
		feedItemInfoEntity    entity.FeedItem
		feedItemInfoDto       dto.FeedItem
	)

	if userId == "" || channelId == "" || itemId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	userListenLaterEntity, err = dao.GetListenLaterByUserIdAndFeedId(ctx, userId, channelId, itemId)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	}
	gconv.Struct(userListenLaterEntity, &userListenLaterDto)

	feedItemInfoEntity, err = dao.GetFeedItemById(ctx, channelId, itemId)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	}
	gconv.Struct(feedItemInfoEntity, &feedItemInfoDto)

	return
}

func GetListenLaterListByUserId(ctx context.Context, userId string, offset, limit int) (userListenLaterDtoList []dto.UserListenLater, err error) {
	var (
		totalCount int
	)

	userListenLaterDtoList, err = dao.GetListenLaterListByUserId(ctx, userId, offset, limit)
	if err != nil {
		return
	}

	totalCount, err = dao.GetTotalListenLaterCountByUserId(ctx, userId)
	if err != nil {
		return
	}

	for i := 0; i < len(userListenLaterDtoList); i++ {
		var dtoItem = userListenLaterDtoList[i]
		if dtoItem.ImageUrl == "" {
			if dtoItem.ChannelImageUrl == "" {
				dtoItem.HasThumbnail = false
			} else {
				dtoItem.ImageUrl = dtoItem.ChannelImageUrl
				dtoItem.HasThumbnail = true
			}
		} else {
			dtoItem.HasThumbnail = true
		}

		if dtoItem.TextDescription == "" && dtoItem.Description != "" {
			rootDocs := soup.HTMLParse(dtoItem.Description)
			dtoItem.TextDescription = rootDocs.FullText()
		}
		if dtoItem.Author == "" {
			dtoItem.Author = dtoItem.ChannelAuthor
		}
		dtoItem.PubDate = formatPubDate(dtoItem.PubDate)
		dtoItem.Duration = formatDuration(dtoItem.Duration)
		dtoItem.Author = formatFeedAuthor(dtoItem.Author)
		dtoItem.Title = formatItemTitle(dtoItem.Title)
		dtoItem.RegDate = consts.ADD_ON_TEXT + formatPubDate(dtoItem.RegDate)
		dtoItem.Count = totalCount
		userListenLaterDtoList[i] = dtoItem
	}

	return
}

func GetListenLaterRSSByUserId(ctx context.Context, userId string) (rss string, err error) {
	var (
		listenLaterDtoList []dto.UserListenLater
		userInfo           dto.UserInfo
		feed               podcast.Podcast
	)

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	userInfo, err = user.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		return
	}
	listenLaterDtoList, err = GetListenLaterListByUserId(ctx, userId, 0, 100)
	if len(listenLaterDtoList) == 0 {
		return
	}

	feedChannelTitle := g.I18n().Tf(ctx, "listen_later_rss_channel_title", userInfo.Nickname)
	feedChannelLink := fmt.Sprintf("https://www.guoshaofm.com/listenlater/%s/rss", userInfo.Id)
	feedChannelDesc := g.I18n().Tf(ctx, "listen_later_rss_channel_description", userInfo.Nickname)
	lastBuildTime := &gtime.NewFromStr(listenLaterDtoList[0].PubDate).Time
	feed = podcast.New(feedChannelTitle, feedChannelLink, feedChannelDesc, lastBuildTime, lastBuildTime)
	feed.Copyright = fmt.Sprintf("Copyright %s GuoshaoFM", userInfo.Nickname)
	feed.AddAuthor(userInfo.Nickname, userInfo.Email)
	feed.AddSummary(feedChannelDesc)
	feed.AddAtomLink("https://www.guoshaofm.com/listenlater/playlist/" + userInfo.Id)
	feed.AddSubTitle(g.I18n().Tf(ctx, "listen_later_rss_channel_description", userInfo.Nickname))
	feed.Generator = "GuoshaoFM (https://www.guoshaofm.com)"
	feed.Language = "zh-CN"
	feed.AddImage("https://www.guoshaofm.com/resource/image/logo192.png")

	for _, listenLaterDtoItem := range listenLaterDtoList {
		var (
			feedItem podcast.Item
		)
		feedItem.AddImage(listenLaterDtoItem.ImageUrl)
		feedItem.AddDuration(gconv.Int64(listenLaterDtoItem.Duration))
		feedItem.AddSummary(listenLaterDtoItem.Description)
		if listenLaterDtoItem.PubDate != "" {
			feedItem.AddPubDate(&gtime.NewFromStr(listenLaterDtoItem.PubDate).Time)
		}
		feedItem.AddEnclosure(listenLaterDtoItem.EnclosureUrl, podcast.MP3, gconv.Int64(listenLaterDtoItem.EnclosureLength))
		feedItem.Title = listenLaterDtoItem.Title
		feedItem.Author = &podcast.Author{
			Name: listenLaterDtoItem.Author,
		}
		feedItem.Description = listenLaterDtoItem.Description
		feedItem.Link = listenLaterDtoItem.Link
		feed.AddItem(feedItem)
	}

	rss = feed.String()

	return
}
