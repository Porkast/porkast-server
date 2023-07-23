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

func UpdateChannelTotalCountJob(ctx context.Context) {
	var (
		err error
	)

	_, err = gcron.Add(ctx, "0 0 */3 * * *", func(ctx context.Context) {
		_ = setChannelTotalCountToCache(ctx)
	}, consts.FEED_CHANNEL_TOTAL_COUNT)

	if err != nil {
		g.Log().Line().Error(ctx, "The ChannelTotalCount job start failed : ", err)
	}
}

func setChannelTotalCountToCache(ctx context.Context) (err error) {
	var (
		totalCount int
	)

	totalCount, err = feed.GetAllFeedChannelCount(ctx)
	if err != nil {
		g.Log().Line().Error(ctx, "Get feed channel total count failed : ", err)
		return
	}

	if totalCount == 0 {
		return
	}
	g.Log().Line().Info(ctx, "The all ZH channel total count is ", totalCount)
	cache.SetCache(ctx, gconv.String(consts.FEED_CHANNEL_TOTAL_COUNT), gconv.String(totalCount), 24*60*60)
	return
}
