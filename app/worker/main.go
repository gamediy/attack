package main

import (
	_ "attack/app/worker/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"

	"attack/app/worker/internal/cmd"
)

func main() {

	cmd.Main.Run(gctx.GetInitCtx())
}
