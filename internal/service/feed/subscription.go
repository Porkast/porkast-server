package feed

import (
	"context"
	"fmt"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao"
	"porkast-server/internal/service/user"

	"github.com/eduncan911/podcast"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

func SubFeedByKeyword(ctx context.Context, userId, keyword, lang, country, excludeFeedId, source string, sortByDate int, subKeywordList []entity.KeywordSubscription) (err error) {

	var (
		userSubKeyword entity.UserSubscription
	)

	userSubKeyword = entity.UserSubscription{
		Id:            guid.S(),
		UserId:        userId,
		Keyword:       keyword,
		Lang:          lang,
		Country:       country,
		ExcludeFeedId: excludeFeedId,
		OrderByDate:   sortByDate,
		Source:        source,
		Status:        1,
		CreateTime:    gtime.Now(),
	}

	err = dao.DoSubKeywordByUserIdAndKeyword(ctx, userSubKeyword, subKeywordList)
	if err != nil {
		return
	}

	return err
}

func GetSubKeywordRSS(ctx context.Context, userId, keyword string) (rssStr string, err error) {

	var (
		userSubKeywordDtoList []dto.UserSubscriptionFeedDetailDto
		userInfo              dto.UserInfo
		feed                  podcast.Podcast
	)

	if userId == "" || keyword == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
	}

	userInfo, err = user.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		return
	}

	userSubKeywordDtoList, err = dao.GetUserSubKeywordListByUserIdAndKeyword(ctx, userId, keyword)
	if err != nil {
		return
	}

	feedChannelTitle := g.I18n().Tf(ctx, "keyword_sub_rss_channel_title", keyword)
	feedChannelLink := fmt.Sprintf("https://www.guoshaofm.com/subscription/%s/%s/rss", userInfo.Id, keyword)
	feedChannelDesc := g.I18n().Tf(ctx, "keyword_sub_rss_channel_description", keyword)
	lastBuildTime := &gtime.NewFromStr(userSubKeywordDtoList[0].CreateDate).Time
	feed = podcast.New(feedChannelTitle, feedChannelLink, feedChannelDesc, lastBuildTime, lastBuildTime)
	feed.Copyright = fmt.Sprintf("Copyright %s GuoshaoFM", userInfo.Nickname)
	feed.AddAuthor(userInfo.Nickname, userInfo.Email)
	feed.AddSummary(feedChannelDesc)
	feed.AddAtomLink("https://www.guoshaofm.com/user/info/" + userInfo.Id)
	feed.AddSubTitle(g.I18n().Tf(ctx, "keyword_sub_rss_channel_description", userInfo.Nickname))
	feed.Generator = "GuoshaoFM (https://www.guoshaofm.com)"
	feed.Language = "zh-CN"
	feed.AddImage("https://www.guoshaofm.com/resource/image/logo192.png")

	for _, userSubKeywordDtoItem := range userSubKeywordDtoList {
		var (
			feedItem podcast.Item
		)
		feedItem.AddImage(userSubKeywordDtoItem.ImageUrl)
		feedItem.AddDuration(gconv.Int64(userSubKeywordDtoItem.Duration))
		feedItem.AddSummary(userSubKeywordDtoItem.Description)
		if userSubKeywordDtoItem.PubDate != "" {
			feedItem.AddPubDate(&gtime.NewFromStr(userSubKeywordDtoItem.PubDate).Time)
		}
		feedItem.AddEnclosure(userSubKeywordDtoItem.EnclosureUrl, podcast.MP3, gconv.Int64(userSubKeywordDtoItem.EnclosureLength))
		feedItem.Title = userSubKeywordDtoItem.Title
		feedItem.Author = &podcast.Author{
			Name: userSubKeywordDtoItem.Author,
		}
		feedItem.Description = userSubKeywordDtoItem.Description
		feedItem.Link = userSubKeywordDtoItem.Link
		feed.AddItem(feedItem)
	}

	rssStr = feed.String()

	return
}

func GetUserSubscriptionCount(ctx context.Context, userId string) (count int, err error) {

	count, err = dao.GetUserSubscriptionCount(ctx, userId)
	if err != nil {
		return
	}

	return
}

func GetUserSubKeywordListByUserId(ctx context.Context, userId string, offset, limit int) (dtoList []dto.UserSubscriptionDto, err error) {

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	if limit == 0 {
		limit = 10
	}

	entities, err := dao.GetUserSubKeywordListByUserId(ctx, userId, offset, limit)
	if err != nil {
		return
	}

	count, err := dao.GetUserSubscriptionCount(ctx, userId)
	if err != nil {
		return
	}

	for _, entity := range entities {
		dtoList = append(dtoList, dto.UserSubscriptionDto{
			Id:          entity.Id,
			UserId:      entity.UserId,
			Keyword:     entity.Keyword,
			Country:     entity.Country,
			OrderByDate: entity.OrderByDate,
			Source:      entity.Source,
			Status:      entity.Status,
			CreateTime:  entity.CreateTime,
			Count:       count,
		})
	}

	return
}

func GetUserSubKeywordRecord(ctx context.Context, userId, keyword, country, excludeFeedId, source string) (result entity.UserSubscription, err error) {

	if userId == "" || keyword == "" || source == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	result, err = dao.GetUserSubKeywordItem(ctx, userId, keyword, country, excludeFeedId, source)

	return
}

func ReactiveUserSubKeyword(ctx context.Context, userId, keyword, country, excludeFeedId string) (err error) {

	err = dao.ActiveUserSubKeywordStatus(ctx, userId, keyword, country, excludeFeedId)

	return
}

func GetItemListByKeywordAndUserId(ctx context.Context, userId, keyword, source string, offset, limit int) (dtoList []dto.FeedItem, err error) {

	if source == "" {
		source = "itunes"
	}

	dtoList, err = dao.GetSubKeywordItemsByUserIdAndKeyword(ctx, userId, keyword, source, offset, limit)
	if err != nil {
		return
	}

	var totalCount = 0
	if len(dtoList) > 0 {
		totalCount, err = dao.GetKeywordSubscriptionCount(ctx, keyword, dtoList[0].Country, source, dtoList[0].ExcludeFeedId)
	}

	for i := 0; i < len(dtoList); i++ {
		dtoList[i].Count = totalCount
	}

	return
}
