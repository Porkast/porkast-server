package ctls

import (
	"context"
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/feed"
	feedService "porkast-server/internal/service/feed"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func (ctl *controller) SubKeyword(req *ghttp.Request) {
	var (
		err          error
		ctx          context.Context
		reqData      *SubSearchKeywordReqData
		ksEntityList []entity.KeywordSubscription
	)

	ctx = req.Context()
	err = req.Parse(&reqData)
	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	subRecord, err := feed.GetUserSubKeywordRecord(ctx, reqData.UserId, reqData.Keyword, reqData.Country, reqData.ExcludeFeedId, reqData.Source)
	if subRecord.Id != "" && subRecord.Status == 1 {
		middleware.JsonExit(req, 1, g.I18n().T(ctx, `{#sub_keyword_exist}`), nil)
	} else if subRecord.Id != "" && subRecord.Status == 0 {
		feed.ReactiveUserSubKeyword(ctx, subRecord.Id, subRecord.Keyword, subRecord.Country, subRecord.ExcludeFeedId)
		middleware.JsonExit(req, 0, g.I18n().T(ctx, `{#sub_keyword_success}`), nil)
	}

	totalSubCount, err := feed.GetUserSubscriptionCount(ctx, reqData.UserId)
	if err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	if totalSubCount >= 3 {
		middleware.JsonExit(req, 1, g.I18n().Tf(ctx, `{#keyword_sub_total_count_limit}`, 10), nil)
	}

	ksEntityList, err = genKeywordSubEntity(ctx, reqData.UserId, reqData.Keyword, reqData.Country, reqData.ExcludeFeedId, reqData.Source)
	if err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	err = feed.SubFeedByKeyword(ctx, reqData.UserId, reqData.Keyword, "", reqData.Country, reqData.ExcludeFeedId, reqData.Source, reqData.SortByDate, ksEntityList)
	if err != nil {
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(ctx, `{#sub_keyword_exist}`), nil)
		}
		g.Log().Line().Error(ctx, err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}

	middleware.JsonExit(req, 0, g.I18n().T(ctx, `{#sub_keyword_success}`), nil)
}

func genKeywordSubEntity(ctx context.Context, userId, keyword, country, excludeFeedId, source string) (ksEntityList []entity.KeywordSubscription, err error) {

	var (
		items       []dto.FeedItem
		searchParam feedService.SearchParams
	)

	searchParam = feedService.SearchParams{
		Keyword:       keyword,
		Page:          0,
		Size:          20,
		ExcludeFeedId: excludeFeedId,
		Country:       country,
	}
	if source == "" || source == "itunes" {
		items, err = feedService.SearchPodcastEpisodesFromItunes(ctx, keyword, country, excludeFeedId)
	} else {
		items, err = feedService.SearchFeedItemsByKeyword(ctx, searchParam)
	}
	if err != nil {
		return nil, err
	}

	err = feedService.BatchStoreFeedItems(ctx, items)
	if err != nil {
		return
	}

	for _, feedItem := range items {

		ksEntity := entity.KeywordSubscription{
			Keyword:       keyword,
			FeedChannelId: feedItem.ChannelId,
			FeedItemId:    feedItem.Id,
			Country:       country,
			ExcludeFeedId: excludeFeedId,
			CreateTime:    gtime.Now(),
			Source:        source,
		}

		ksEntityList = append(ksEntityList, ksEntity)
	}

	return
}

func (ctl *controller) GetSubKeywordFeedRSS(req *ghttp.Request) {

	var (
		err     error
		ctx     context.Context
		userId  string
		keyword string
		rssStr  string
	)

	userId = req.Get("userId").String()
	keyword = req.Get("keyword").String()

	rssStr, err = feed.GetSubKeywordRSS(ctx, userId, keyword)
	if err != nil {
		// TODO: Add error page
		g.Log().Line().Error(ctx, err)
	}

	req.Response.WriteXml(rssStr)
}


func (ctl *controller) GetUserSubKeywordListAPI(req *ghttp.Request) {
	var (
		err      error
		ctx      context.Context
		userId   string
		page     int
		offset   int
		limit    = 10
		itemList []dto.UserSubKeywordDto
	)

	ctx = req.Context()
	userId = req.Get("userId").String()
	page = req.Get("page").Int()
	if page == 0 {
		page = 1
	}

	offset = (page - 1) * limit
	itemList, err = feed.GetUserSubKeywordListByUserId(ctx, userId, offset, limit)
	if err != nil {
		middleware.JsonExit(req, 1, err.Error())
	} else {
		middleware.JsonExit(req, 0, "success", itemList)
	}

}

func (ctl *controller) GetSubKeywordItemListAPI(req *ghttp.Request) {
	var (
		err      error
		ctx      context.Context
		userId   string
		keyword  string
		source   string
		page     int
		offset   int
		limit    = 10
		itemList []dto.FeedItem
	)
	ctx = req.Context()
	userId = req.Get("userId").String()
	page = req.Get("page").Int()
	keyword = req.Get("keyword").String()
	source = req.Get("source").String()
	if page == 0 {
		page = 1
	}

	offset = (page - 1) * limit

	itemList, err = feed.GetItemListByKeywordAndUserId(ctx, userId, keyword, source, offset, limit)
	if err != nil {
		middleware.JsonExit(req, 1, err.Error())
	} else {
		middleware.JsonExit(req, 0, "success", itemList)
	}
}
