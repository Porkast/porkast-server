package jobs

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/cache"
	"guoshao-fm-web/internal/service/internal/dao"
	"sync"
	"sync/atomic"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/grpool"
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
				atomic.AddInt64(&totalCount, gconv.Int64(count))
			}
		})
	}

	wg.Wait()

	cache.SetCache(ctx, gconv.String(consts.FEED_ITEM_TOTAL_COUNT), gconv.String(totalCount), 0)
	return
}
