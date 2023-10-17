package utility

import (
	"context"

	"github.com/mmcdole/gofeed"
)

func ParseFeed(ctx context.Context, rssStr string) (feed *gofeed.Feed) {

	var (
		err error
		fp  *gofeed.Parser
	)
	fp = gofeed.NewParser()
	feed, err = fp.ParseString(rssStr)
	if err != nil {
		return nil
	}
	return
}

func IsStringRSSXml(respStr string) bool {
	var (
		err  error
		fp   *gofeed.Parser
		feed *gofeed.Feed
	)
	if respStr != "" {
		fp = gofeed.NewParser()
		feed, err = fp.ParseString(respStr)
		if err != nil || feed == nil {
			return false
		}
		return true
	}

	return false
}
