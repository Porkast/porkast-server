package gslog

import (
	"context"
	"porkast-server/internal/service/elasticsearch"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

func Init() {

	var LoggingHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {

		go func(ctx context.Context, in *glog.HandlerInput) {
			time := in.TimeFormat
			level := gstr.Trim(in.LevelFormat, "[]")
			content := gstr.Trim(in.Content)
			elasticsearch.GetClient().StoreLogs(ctx, time, level, content)
		}(ctx, in)

		in.Next(ctx)
	}

	g.Log().SetHandlers(LoggingHandler)
}
