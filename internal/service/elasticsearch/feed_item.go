package elasticsearch

import (
	"context"
	"guoshao-fm-web/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

func (c *GSElastic) QueryFeedItemFull(ctx context.Context, keyword string, sortByDate, from, size int) (esFeedItemList []entity.FeedItemESData, err error) {
	simpleStringQuery := elastic.NewSimpleQueryStringQuery(keyword)
	simpleStringQuery.FieldWithBoost("title", 10)
	simpleStringQuery.FieldWithBoost("textDescription", 2)
	simpleStringQuery.FieldWithBoost("author", 1)
	// simpleStringQuery.FieldWithBoost("channelTitle", 1)
    simpleStringQuery.MinimumShouldMatch("75%")

    termQuery := elastic.NewTermQuery("language", "zh-CN")
    termQuery.CaseInsensitive(true)

	highlight := elastic.NewHighlight()
	highlight = highlight.PreTags("<span style='color: red;'>").PostTags("</span>")
	highlight = highlight.Fields(elastic.NewHighlighterField("title"), elastic.NewHighlighterField("textDescription"), elastic.NewHighlighterField("author"))
	searchService := c.Client.Search().
		Index("feed_item").
		Query(simpleStringQuery).
        PostFilter(termQuery).
		Highlight(highlight).
		From(from).Size(size).
		Pretty(true)

	if sortByDate == 1 {
		searchService.Sort("pubDate", false)
	}

	searchResult, err := searchService.Do(ctx)
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
