package cmd

import (
	"context"
	"os"

	"guoshao-fm-web/internal/routers"
	"guoshao-fm-web/internal/service/cache"
	"guoshao-fm-web/internal/service/elasticsearch"

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
}
