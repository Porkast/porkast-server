package feed

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func formatPubDate(pubDate string) (formatPubDate string) {
	formatPubDate = gtime.New(pubDate).Format("Y-m-d")
	return
}

func formatDuration(duration string) (formatDuration string) {
	if !gstr.Contains(duration, ":") {
		var (
			totalSecs = gconv.Int(duration)
			hours     int
			minutes   int
			seconds   int
		)
		hours = totalSecs / 3600
		minutes = (totalSecs % 3600) / 60
		seconds = totalSecs % 60
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
