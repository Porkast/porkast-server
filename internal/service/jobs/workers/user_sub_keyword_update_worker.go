package workers

import (
	"porkast-server/internal/consts"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/elasticsearch"
	"porkast-server/internal/service/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

func UpdateUserSubkeyword(keyword, lang string, orderByDate int) {
	var (
		err                error
		ctx                = gctx.New()
		esFeedItemDataList []entity.FeedItemESData
	)

	esFeedItemDataList, err = elasticsearch.GetClient().QueryFeedItemFull(ctx, keyword, orderByDate, 0, 20)
	if err != nil {
		g.Log().Line().Errorf(ctx, "search by keyword %s , orderByDate %d failed:\n%s", keyword, orderByDate, err)
		return
	}

	for _, esFeedItem := range esFeedItemDataList {
		var (
			keywordSubEntity entity.KeywordSubscription
		)

		keywordSubEntity = entity.KeywordSubscription{
			Id:            guid.S(),
			Keyword:       keyword,
			FeedChannelId: esFeedItem.ChannelId,
			FeedItemId:    esFeedItem.Id,
			CreateTime:    gtime.Now(),
			OrderByDate:   orderByDate,
			Lang:          lang,
		}

		err = dao.CreateKeywordSubScriptionEntity(ctx, keywordSubEntity)
		if err != nil {
			if err.Error() == consts.DB_DATA_ALREADY_EXIST {
				g.Log().Line().Debugf(ctx, "keywordSubEntity already exist, ignore.")
			} else {
				g.Log().Line().Errorf(ctx, "create keywordSubEntity failed:\n%s", err)
			}
		}
	}

}
