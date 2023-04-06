package main

import (
	_ "guoshao-fm-web/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"guoshao-fm-web/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
