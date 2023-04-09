package elasticsearch

import (
	"context"
	"guoshao-fm-web/internal/model/entity"
	"reflect"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

func (c *GSElastic) QueryFeedItemFull(ctx context.Context, keyword string, from, size int) (esFeedItemList []entity.FeedItemESData, err error) {
	multMatch := elastic.NewMultiMatchQuery(keyword, "title", "author", "description")
	multMatch.FieldWithBoost("title", 3)
	multMatch.FieldWithBoost("author", 2)
	multMatch.FieldWithBoost("description", 1)
	searchResult, err := c.Client.Search().
		Index("feed_item").
		Query(multMatch).
		From(from).Size(size).
		Pretty(true).
		Do(ctx)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	}

	var itemType entity.FeedItemESData
	var totalCount int
	totalCount = int(searchResult.TotalHits())
	for _, item := range searchResult.Each(reflect.TypeOf(itemType)) {
		if t, ok := item.(entity.FeedItemESData); ok {
			t.Count = totalCount
			esFeedItemList = append(esFeedItemList, t)
		}
	}

	return
}
