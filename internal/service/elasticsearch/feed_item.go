package elasticsearch

import (
	"context"
	"guoshao-fm-web/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

func (c *GSElastic) QueryFeedItemFull(ctx context.Context, keyword string, from, size int) (esFeedItemList []entity.FeedItemESData, err error) {
	multMatch := elastic.NewMultiMatchQuery(keyword, "title", "author", "textDescription")
	multMatch.FieldWithBoost("title", 10)
	multMatch.FieldWithBoost("textDescription", 2)
	multMatch.FieldWithBoost("author", 1)
	multMatch.FieldWithBoost("channelTitle", 1)
    multMatch.Type("most_fields")
	highlight := elastic.NewHighlight()
	highlight = highlight.PreTags("<span style='color: red;'>").PostTags("</span>")
	highlight = highlight.Fields(elastic.NewHighlighterField("title"), elastic.NewHighlighterField("channelTitle"), elastic.NewHighlighterField("textDescription"), elastic.NewHighlighterField("author"))
	searchResult, err := c.Client.Search().
		Index("feed_item").
		Query(multMatch).
		Highlight(highlight).
		From(from).Size(size).
		Pretty(true).
		Do(ctx)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	}

	var totalCount int
	var tookTime float64
	totalCount = int(searchResult.TotalHits())
	tookTime = float64(searchResult.TookInMillis) / float64(time.Microsecond)
	g.Log().Line().Debug(ctx, "took time :", searchResult.TookInMillis)
	for _, hit := range searchResult.Hits.Hits {
		var esFeedItem entity.FeedItemESData
		gjson.Unmarshal(hit.Source, &esFeedItem)

		esFeedItem.Count = totalCount
		esFeedItem.TookTime = tookTime
		if len(hit.Highlight["title"]) != 0 {
			esFeedItem.HighlightTitle = hit.Highlight["title"][0]
		}
		if len(hit.Highlight["channelTitle"]) != 0 {
			esFeedItem.ChannelTitle = hit.Highlight["channelTitle"][0]
		}
		if len(hit.Highlight["textDescription"]) != 0 {
			esFeedItem.TextDescription = hit.Highlight["textDescription"][0]
		}
		if len(hit.Highlight["author"]) != 0 {
			esFeedItem.Author = hit.Highlight["author"][0]
		}
		esFeedItemList = append(esFeedItemList, esFeedItem)
	}

	return
}
