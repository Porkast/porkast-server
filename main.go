package main

import (
	_ "porkast-server/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"porkast-server/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
