package cmd

import (
	"context"
	"os"

	"guoshao-fm-web/internal/routers"

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
			s := g.Server()
			s.Group("/", routers.WebRouter)
			initConfig()
			s.Run()
			return
		},
	}
)

func initConfig() {
	if os.Getenv("env") == "dev" {
		genv.Set("GF_GCFG_FILE", "config.dev.yaml")
	}
}
