package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

func GetFeedItemByItemId(ctx context.Context, itemId string) (feedItemInfoDto dto.FeedItem, err error) {
	var (
		feedItemModel entity.FeedItem
	)

	feedItemModel, err = dao.GetFeedItemById(ctx, itemId)
	gconv.Struct(feedItemModel, &feedItemModel)

	return
}
