package ctls

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/feed"
	feedService "guoshao-fm-web/internal/service/feed"
	"guoshao-fm-web/internal/service/middleware"

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

	totalSubCount, err := feed.GetUserSubscriptionCount(ctx, reqData.UserId)
	if err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	if totalSubCount >= 10 {
		middleware.JsonExit(req, 1, g.I18n().Tf(ctx, `{#keyword_sub_total_count_limit}`, 10), nil)
	}

	ksEntityList, err = genKeywordSubEntity(ctx, reqData.UserId, reqData.Keyword, reqData.Lang, reqData.SortByDate)
	if err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	err = feed.SubFeedByKeyword(ctx, reqData.UserId, reqData.Keyword, reqData.Lang, reqData.SortByDate, ksEntityList)
	if err != nil {
		if err.Error() == consts.DB_DATA_ALREADY_EXIST {
			middleware.JsonExit(req, 1, g.I18n().T(ctx, `{#sub_keyword_exist}`), nil)
		}
		g.Log().Line().Error(ctx, err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}

	middleware.JsonExit(req, 0, g.I18n().T(ctx, `{#sub_keyword_success}`), nil)
}

func genKeywordSubEntity(ctx context.Context, userId, keyword, lang string, sortByDate int) (ksEntityList []entity.KeywordSubscription, err error) {

	var (
		items       []dto.FeedItem
		searchParam feedService.SearchParams
	)

	searchParam = feedService.SearchParams{
		Keyword:    keyword,
		Page:       0,
		Size:       20,
		SortByDate: sortByDate,
	}
	items, err = feedService.SearchFeedItemsByKeyword(ctx, searchParam)
	if err != nil {
		return nil, err
	}

	for _, feedItem := range items {
		ksEntity := entity.KeywordSubscription{
			Id:            userId,
			Keyword:       keyword,
			FeedChannelId: feedItem.ChannelId,
			FeedItemId:    feedItem.Id,
			Lang:          lang,
			OrderByDate:   sortByDate,
			CreateTime:    gtime.Now(),
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
