package feed

import (
	"context"
	"fmt"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"
	"guoshao-fm-web/internal/service/user"

	"github.com/GuoShaoOrg/feeds"
	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

func CreateListenLaterByUserIdAndFeedId(ctx context.Context, userId, channelId, itemId string) (err error) {
	var (
		newEntity entity.UserListenLater
	)

	newEntity = entity.UserListenLater{
		Id:        guid.S(),
		UserId:    userId,
		ChannelId: channelId,
		ItemId:    itemId,
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
		feed               feeds.Rss
		userInfo           dto.UserInfo
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

	feed.Title = g.I18n().Tf(ctx, "listen_later_rss_channel_title", userInfo.Nickname)
	feed.Author = &feeds.Author{
		Name: g.I18n().Tf(ctx, "listen_later_rss_channel_auther", userInfo.Nickname),
	}
	if userInfo.Email != "" {
		feed.Author.Email = userInfo.Email
	}
	feed.Description = g.I18n().Tf(ctx, "listen_later_rss_channel_description", userInfo.Nickname)
	feed.Link = &feeds.Link{
		Href: fmt.Sprintf("https://www.guoshaofm.com/listenlater/%s/rss", userInfo.Id),
		Rel:  "self",
		Type: "application/rss+xml",
	}

    //TODO: add listen later images
    feed.Image = &feeds.Image{
    	Url:    "",
    	Title:  "",
    	Link:   "",
    	Width:  0,
    	Height: 0,
    }


	feed.Items = make([]*feeds.Item, 0)
	for _, listenLaterDtoItem := range listenLaterDtoList {
		var (
			feedItem *feeds.Item
		)
		feedItem = &feeds.Item{
			Title: listenLaterDtoItem.Title,
			Link: &feeds.Link{
				Href: listenLaterDtoItem.Link,
				Rel:  "self",
				Type: "application/rss+xml",
			},
			Source: &feeds.Link{
				Href: listenLaterDtoItem.Link,
				Rel:  "self",
				Type: "application/rss+xml",
			},
			Author: &feeds.Author{
				Name: listenLaterDtoItem.Author,
			},
			Description: listenLaterDtoItem.Description,
			Id:          listenLaterDtoItem.ItemId,
			Updated:     gtime.NewFromStr(listenLaterDtoItem.PubDate).Time,
			Created:     gtime.NewFromStr(listenLaterDtoItem.PubDate).Time,
			Enclosure: &feeds.Enclosure{
				Url:    listenLaterDtoItem.EnclosureUrl,
				Length: listenLaterDtoItem.EnclosureLength,
				Type:   listenLaterDtoItem.EnclosureType,
			},
			Content: listenLaterDtoItem.Description,
		}

        feed.Items = append(feed.Items, feedItem)
	}

    rss, err = feed.ToAtom()

	return
}
