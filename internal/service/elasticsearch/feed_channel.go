package elasticsearch

import (
	"context"
	"guoshao-fm-web/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

func (c *GSElastic) QueryFeedChannelFull(ctx context.Context, keyword string, offset, limit int) (esFeedChannelList []entity.FeedChannelESData, err error) {
	simpleStringQuery := elastic.NewSimpleQueryStringQuery(keyword)
	simpleStringQuery.FieldWithBoost("title", 4)
	simpleStringQuery.FieldWithBoost("author", 1)
	simpleStringQuery.MinimumShouldMatch("75%")
	highlight := elastic.NewHighlight()
	highlight = highlight.PreTags("<span style='color: red;'>").PostTags("</span>")
	highlight = highlight.Fields(
		elastic.NewHighlighterField("title"),
		elastic.NewHighlighterField("author"),
	)
	searchResult, err := c.Client.Search().
		Index("feed_channel").
		Query(simpleStringQuery).
		Highlight(highlight).
		From(offset).Size(limit).
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
		var esFeedChannel entity.FeedChannelESData
		gjson.Unmarshal(hit.Source, &esFeedChannel)

		esFeedChannel.Count = totalCount
		esFeedChannel.TookTime = tookTime
		if len(hit.Highlight["title"]) != 0 {
			esFeedChannel.Title = hit.Highlight["title"][0]
		}
		if len(hit.Highlight["author"]) != 0 {
			esFeedChannel.Author = hit.Highlight["author"][0]
		}
		esFeedChannelList = append(esFeedChannelList, esFeedChannel)
	}

	return
}
