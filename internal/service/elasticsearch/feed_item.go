package elasticsearch

import (
	"context"
	"guoshao-fm-web/internal/model/entity"
	"reflect"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

func (c *GSElastic) QueryFeedItemFull(ctx context.Context, keyword string) (esFeedItemList []entity.FeedItemESData) {
	multMatch := elastic.NewMultiMatchQuery(keyword, "title", "author", "description")
	multMatch.FieldWithBoost("title", 3)
	multMatch.FieldWithBoost("author", 2)
	multMatch.FieldWithBoost("description", 1)
	searchResult, err := c.Client.Search().
		Index("feed_item").
		Query(multMatch).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		g.Log().Line().Error(ctx, err)
	}

	var itemType entity.FeedItemESData
	for _, item := range searchResult.Each(reflect.TypeOf(itemType)) {
		if t, ok := item.(entity.FeedItemESData); ok {
			esFeedItemList = append(esFeedItemList, t)
		}
	}

	return
}
