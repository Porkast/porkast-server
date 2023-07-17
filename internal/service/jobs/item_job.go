package jobs

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/cache"
	"guoshao-fm-web/internal/service/internal/dao"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func UpdateItemTotalCountJob(ctx context.Context) {
	var (
		err error
	)

	_, err = gcron.Add(ctx, "0 0 2 * * *", func(ctx context.Context) {
		_ = setItemTotalCountToCache(ctx)
	}, consts.FEED_ITEM_TOTAL_COUNT)

	if err != nil {
		g.Log().Line().Error(ctx, "The ChannelTotalCount job start failed : ", err)
	}
}

func setItemTotalCountToCache(ctx context.Context) (err error) {
	var (
		totalCount      int64
		feedChannelList []entity.FeedChannel
	)

	wg := sync.WaitGroup{}
	pool := grpool.New(100)
	feedChannelList, err = dao.GetZHFeedChannelList(ctx)
	for _, feedChannel := range feedChannelList {
		feedChannelTemp := feedChannel
		wg.Add(1)
		pool.Add(ctx, func(ctx context.Context) {
			defer wg.Done()
			count, err := dao.GetFeedItemCountByChannelId(ctx, feedChannelTemp.Id)
			if err == nil {
				g.Log().Line().Infof(ctx, "The channel %s item total count is %d", feedChannelTemp.Title, count)
				atomic.AddInt64(&totalCount, gconv.Int64(count))
			}
		})
	}

	wg.Wait()

	if totalCount == 0 {
		return
	}
	g.Log().Line().Infof(ctx, "The all ZH items total count is %d", totalCount)
	cache.SetCache(ctx, gconv.String(consts.FEED_ITEM_TOTAL_COUNT), gconv.String(totalCount), int(24*60*60))
	return
}

func UpdateLatestItemListCountJob(ctx context.Context) {
	var (
		err error
	)

	_, err = gcron.Add(ctx, "0 0 2 * * *", func(ctx context.Context) {
		_ = setLatestFeedItems(ctx)
	}, consts.TODAY_FEED_ITEM_LIST)

	if err != nil {
		g.Log().Line().Error(ctx, "The UpdateItemTotalCountJob job start failed : ", err)
	}
}

func setLatestFeedItems(ctx context.Context) (err error) {
	var (
		startDate    *gtime.Time
		startDateStr string
		endDate      *gtime.Time
		endDateStr   string
		itemList     []dto.FeedItem
		itemListJson *gjson.Json
	)

	startDate = gtime.Now().StartOfDay()
	endDate = gtime.Now().EndOfDay()

	startDateStr = startDate.ISO8601()
	endDateStr = endDate.ISO8601()

	itemList = dao.GetFeedItemListByPubDate(ctx, startDateStr, endDateStr)
	if err != nil {
		g.Log().Line().Error(ctx, "Get latest feed items failed: ", err)
		return
	}

	if len(itemList) == 0 {
		return
	}

	itemListJson = gjson.New(itemList)
	if err != nil {
		g.Log().Line().Error(ctx, "Decode feed item list to json failed", err)
		return
	}
	cache.SetCache(ctx, gconv.String(consts.TODAY_FEED_ITEM_LIST), itemListJson.MustToJsonString(), int(time.Second*60*60))

	return
}
