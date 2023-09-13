package feed

import (
	"context"
	"fmt"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"
	"guoshao-fm-web/internal/service/user"

	"github.com/eduncan911/podcast"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

func SubFeedByKeyword(ctx context.Context, userId, keyword, lang string, sortByDate int, subKeywordList []entity.KeywordSubscription) (err error) {

	var (
		userSubKeyword entity.UserSubKeyword
	)

	userSubKeyword = entity.UserSubKeyword{
		Id:          guid.S(),
		UserId:      userId,
		Keyword:     keyword,
		Lang:        lang,
		OrderByDate: sortByDate,
		Status:      1,
		CreateTime:  gtime.Now(),
	}

	err = dao.DoSubKeywordByUserIdAndKeyword(ctx, userSubKeyword, subKeywordList)
	if err != nil {
		return
	}

	return err
}

func GetSubKeywordRSS(ctx context.Context, userId, keyword string) (rssStr string, err error) {

	var (
		userSubKeywordDtoList []dto.UserSubKeyword
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
