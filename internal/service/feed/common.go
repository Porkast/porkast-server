package feed

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gtime"
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
		}
	}
	return
}
