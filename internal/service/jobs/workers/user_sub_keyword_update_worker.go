package workers

import (
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/feed"
	"porkast-server/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func UpdateUserSubkeyword(keyword, country, excludeFeedId, source string) {
	var (
		err          error
		ctx          = gctx.New()
		feedItemList []dto.FeedItem
	)

	if source == "" || source == "itunes" {
		feedItemList, err = feed.SearchPodcastEpisodesFromItunes(ctx, keyword, country, excludeFeedId)
		if err != nil {
			g.Log().Line().Errorf(ctx, "search by keyword %s , excludeFeedId %s failed:\n%s", keyword, excludeFeedId, err)
			return
		}
	}

	if err != nil {
		g.Log().Line().Errorf(ctx, "search by keyword %s , excludeFeedId %s failed:\n%s", keyword, excludeFeedId, err)
		return
	}

	for _, feedItem := range feedItemList {
		var (
			keywordSubEntity entity.KeywordSubscription
		)

		keywordSubEntity = entity.KeywordSubscription{
			Keyword:       keyword,
			FeedChannelId: feedItem.ChannelId,
			FeedItemId:    feedItem.Id,
			CreateTime:    gtime.Now(),
			ExcludeFeedId: excludeFeedId,
			Country:       country,
			Source:        source,
		}

		err = dao.CreateKeywordSubScriptionEntity(ctx, keywordSubEntity)
		if err != nil {
			if err.Error() == consts.DB_DATA_ALREADY_EXIST {
				g.Log().Line().Debugf(ctx, "keywordSubEntity already exist, ignore.")
			} else {
				g.Log().Line().Errorf(ctx, "create keywordSubEntity failed:\n%s", err)
			}
			return
		}

		batchItems := make([]dto.FeedItem, 0)
		batchItems = append(batchItems, feedItem)
		err = feed.BatchStoreFeedItems(ctx, batchItems)
		if err != nil {
			g.Log().Line().Errorf(ctx, "batch store feed items failed:\n%s", err)
			return
		}
	}

}
