package ctls

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	feedService "guoshao-fm-web/internal/service/feed"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) SearchResult(req *ghttp.Request) {
	var (
		err           error
		searchKeyword string
		scope         string
		page          int
		sortByDate    int
		searchParam   feedService.SearchParams
		tplMap        map[string]interface{}
	)

	searchKeyword = req.GetQuery("q").String()
	scope = req.GetQuery("scope").String()
	page = req.GetQuery("page").Int()
	sortByDate = req.GetQuery("sortByDate").Int()

	searchParam = feedService.SearchParams{
		Keyword:    searchKeyword,
		Page:       page,
		Scope:      scope,
		SortByDate: sortByDate,
	}

	if scope == consts.SEARCH_CHANNEL_SCOPE {
		tplMap, err = searchFeedChannels(req.GetCtx(), searchParam)
	} else {
		tplMap, err = searchFeedItems(req.GetCtx(), searchParam)
	}

	if err != nil {
		//TODO: Add error page
	}

	req.Response.WriteTpl("search.html", tplMap)
}

func searchFeedItems(ctx context.Context, searchParam feedService.SearchParams) (map[string]interface{}, error) {
	var (
		err                  error
		totalPage            int
		totalCount           int
		totalCountText       string
		tookTime             float64
		tookTimeStr          string
		tookTimeText         string
		items                []dto.FeedItem
		channels             []dto.FeedChannel
		subscriptionBtnText  string
		subConfirmModalTitle string
		subConfirmModalDesc  string
	)

	if searchParam.Page == 0 || searchParam.Page == 1 {
		channels, err = feedService.QueryFeedChannelByKeyword(ctx, searchParam)
		if err != nil {
			return nil, err
		}
	}

	items, err = feedService.SearchFeedItemsByKeyword(ctx, searchParam)
	if err != nil {
		return nil, err
	}
	if len(items) > 0 {
		tookTime = items[0].TookTime
		tookTimeStr = strconv.FormatFloat(tookTime, 'f', -3, 64)
		totalCount = items[0].Count
		totalPage = totalCount / 10
		if totalCount%10 > 0 {
			totalPage = totalPage + 1
		}
	}

	totalCountText = g.I18n().Tf(ctx, consts.SEARCH_RESULT_COUNT_TEXT_VALUE, totalCount)
	tookTimeText = g.I18n().Tf(ctx, consts.SEARCH_TOOK_TIME_TEXT_VALUE, tookTimeStr)
	subscriptionBtnText = g.I18n().Tf(ctx, "keyword_sub_btn_text", searchParam.Keyword)
	subConfirmModalTitle = g.I18n().Tf(ctx, "sub_confirm_modal_title")
	if searchParam.SortByDate == 1 {
		subConfirmModalDesc = g.I18n().Tf(ctx, "sub_confirm_modal_desc_order_by_date", searchParam.Keyword)
	} else {
		subConfirmModalDesc = g.I18n().Tf(ctx, "sub_confirm_modal_desc_relative", searchParam.Keyword)
	}
	var tplMap = consts.GetCommonTplMap(ctx)
	tplMap[consts.SEARCH_KEYWORD] = searchParam.Keyword
	tplMap[consts.CURRENT_PAGE] = searchParam.Page
	tplMap[consts.TOTAL_PAGE] = totalPage
	tplMap[consts.SEARCH_RESULT_COUNT_TEXT] = totalCountText
	tplMap[consts.SEARCH_TOOK_TIME_TEXT] = tookTimeText
	tplMap[consts.SUB_KEYWORD_BTN_TEXT] = subscriptionBtnText
	tplMap[consts.SUB_CONFIRM_MODAL_TITLE] = subConfirmModalTitle
	tplMap[consts.SUB_CONFIRM_MODAL_DESC] = subConfirmModalDesc
	tplMap[consts.FEED_ITEMS] = items
	tplMap[consts.FEED_CHANNELS] = channels
	if searchParam.SortByDate == 1 {
		tplMap[consts.SEARCH_ORDER_BY_DATE] = true
	}

	return tplMap, nil
}

func searchFeedChannels(ctx context.Context, searchParam feedService.SearchParams) (map[string]interface{}, error) {
	var (
		err            error
		totalPage      int
		totalCount     int
		totalCountText string
		tookTime       float64
		tookTimeStr    string
		tookTimeText   string
		channels       []dto.FeedChannel
		tplMap         = consts.GetCommonTplMap(ctx)
	)

	channels, err = feedService.QueryFeedChannelByKeyword(ctx, searchParam)
	if err != nil {
		return nil, err
	}

	if len(channels) > 0 {
		tookTime = channels[0].TookTime
		tookTimeStr = strconv.FormatFloat(tookTime, 'f', -3, 64)
		totalCount = channels[0].Count
		totalPage = totalCount / 10
		if totalCount%10 > 0 {
			totalPage = totalPage + 1
		}
	}

	totalCountText = g.I18n().Tf(ctx, consts.SEARCH_RESULT_COUNT_TEXT_VALUE, totalCount)
	tookTimeText = g.I18n().Tf(ctx, consts.SEARCH_TOOK_TIME_TEXT_VALUE, tookTimeStr)
	tplMap[consts.SEARCH_KEYWORD] = searchParam.Keyword
	tplMap[consts.CURRENT_PAGE] = searchParam.Page
	tplMap[consts.TOTAL_PAGE] = totalPage
	tplMap[consts.SEARCH_RESULT_COUNT_TEXT] = totalCountText
	tplMap[consts.SEARCH_TOOK_TIME_TEXT] = tookTimeText
	tplMap[consts.FEED_CHANNELS] = channels
	tplMap[consts.SEARCH_CHANNEL] = true
	tplMap[consts.SEARCH_ONLY_MATCH_TITLE] = g.I18n().T(ctx, consts.SEARCH_ONLY_MATCH_TITLE)
	if searchParam.SortByDate == 1 {
		tplMap[consts.SEARCH_ORDER_BY_DATE] = true
	}

	return tplMap, nil
}
