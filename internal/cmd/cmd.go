package cmd

import (
	"context"
	"os"

	"porkast-server/internal/consts"
	"porkast-server/internal/routers"
	"porkast-server/internal/service/cache"
	"porkast-server/internal/service/celery"
	"porkast-server/internal/service/elasticsearch"
	"porkast-server/internal/service/gslog"
	"porkast-server/internal/service/jobs"
	"porkast-server/internal/service/jobs/workers"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
)

var (
	Main = gcmd.Command{
		Name:  "Guoshao FM Web",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initConfig()
			s := g.Server()
			s.Group("/", routers.WebRouter)
			s.Group("/v1/api", routers.V1ApiRouter)
			initComponent(ctx)
			s.Run()
			return
		},
	}
)

func initConfig() {
	if os.Getenv("env") == "dev" {
		genv.Set("GF_GCFG_FILE", "config.dev.yaml")
	} else if os.Getenv("env") == "prod" {
		genv.Set("GF_GCFG_FILE", "config.prod.yaml")
	} else {
		genv.Set("GF_GCFG_FILE", "config.yaml")
	}
	g.I18n().SetPath("./resource/i18n")
}

func initComponent(ctx context.Context) {
	cache.InitCache(ctx)
	elasticsearch.InitES(ctx)
	gslog.Init()
	jobs.InitJobs(ctx)
	celery.InitCeleryClient(ctx)
	registerCeleryJobs(ctx)
	celery.GetClient().StartWorker()

	jobs.UpdateUserSubKeywordJobs(ctx)
}

func registerCeleryJobs(ctx context.Context) {
	celery.GetClient().Register(consts.USER_SUB_KEYWORD_UPDATE, workers.UpdateUserSubkeyword)
}
