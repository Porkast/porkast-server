package main

import (
	_ "guoshao-fm-web/internal/packed"

	"guoshao-fm-web/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
