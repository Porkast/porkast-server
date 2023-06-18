package jobs

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/service/cache"
	"guoshao-fm-web/internal/service/feed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/util/gconv"
)

func SetItemTotalCountJob(ctx context.Context) {
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
		totalCount int
	)

	totalCount, err = feed.GetAllFeedItemCount(ctx)
	if err != nil {
		g.Log().Line().Error(ctx, "Get feed item total count failed : ", err)
		return
	}

	cache.SetCache(ctx, gconv.String(consts.FEED_ITEM_TOTAL_COUNT), gconv.String(totalCount), 0)
	return
}
