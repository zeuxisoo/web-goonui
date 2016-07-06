package routes

import (
    "github.com/zeuxisoo/go-goonui/modules/context"
)

func Home(ctx *context.Context) {
    ctx.HTML(200, "index")
}
