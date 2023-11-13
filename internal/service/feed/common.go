package feed

import (
	"context"
	"fmt"
	"strconv"

	"github.com/anaskhan96/soup"
	"github.com/araddon/dateparse"
	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mmcdole/gofeed"
)

func formatPubDate(pubDate string) (formatPubDate string) {
	t, err := dateparse.ParseLocal(pubDate)
	if err != nil {
		return
	}

	formatPubDate = t.Format("2006-01-02")

	return
}

func formatDuration(duration string) (formatDuration string) {
	if !gstr.Contains(duration, ":") {
		var (
			totalMillSecs = gconv.Int(duration)
			hours         int
			minutes       int
			seconds       int
		)
		hours = totalMillSecs / 3600000
		minutes = (totalMillSecs % 3600000) / 60000
		seconds = (totalMillSecs % 60000) / 1000
		formatDuration = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	} else {
		var (
			splits []string
		)
		splits = gstr.Split(duration, ":")
		if len(splits) < 3 {
			formatDuration = "00:" + duration
		} else {
			formatDuration = duration
		}
	}
	return
}

func formatSourceLink(link string) (formatLink string) {
	if gstr.Contains(link, "ximalaya.com//") {
		formatLink = gstr.Replace(link, "ximalaya.com//", "ximalaya.com/")
	} else {
		formatLink = link
	}

	return
}

func formatItemShownotes(shownots string) (formatShownotes string) {
	var (
		matches [][]string
		err     error
	)

	matches, err = gregex.MatchAllString(`((\d\d):([0-5][0-9]):([0-5]\d))|([0-5][0-9]):([0-5]\d)`, shownots)
	if err != nil {
		g.Log().Line().Debug(gctx.New(), err)
		return shownots
	}

	formatShownotes = shownots
	for _, match := range matches {
		var matchItem = match[0]
		formatShownotes = gstr.Replace(formatShownotes, matchItem, `<span class='underline hover:cursor-pointer' onclick='playAt("`+matchItem+`")'>`+matchItem+`</span>`)
	}

	return
}

func formatFeedAuthor(author string) (formatAuthor string) {

	if author != "" && gstr.HasSuffix(author, "|") {
		formatAuthor = author[:len(author)-1]
	} else {
		formatAuthor = author
	}

	return
}

func formatTitle(title string) (formatTitle string) {

	formatTitle = title
	if title != "" {
		docs := soup.HTMLParse(title)
		if docs.Error == nil {
			formatTitle = docs.FullText()
		}
	}

	return
}

func formatItemTitle(title string) (formatTitle string) {

	if gstr.Contains(title, "\"") {
		formatTitle = gstr.Replace(title, "\"", "`")
	} else {
		formatTitle = title
	}

	return
}

func GenerateFeedItemId(feedUrl, trackName string) (itemID string) {
	itemID = strconv.FormatUint(ghash.RS64([]byte(feedUrl+trackName)), 32)
	return
}

func GenerateFeedChannelId(feedUrl, collectionName string) (channelId string) {
	channelId = strconv.FormatUint(ghash.RS64([]byte(feedUrl+collectionName)), 32)
	return
}

func GeneratePlaylistId(name, userId string) (itemID string) {
	itemID = strconv.FormatUint(ghash.RS64([]byte(name+userId)), 32)
	return
}

func GeneratePlaylistItemId(playlistId, itemId string) (itemID string) {
	itemID = strconv.FormatUint(ghash.RS64([]byte(playlistId+itemId)), 32)
	return
}

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
