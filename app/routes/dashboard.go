package routes

import (
    "github.com/zeuxisoo/go-goonui/app/kernels/context"
)

func Dashboard(ctx *context.Context) {
    ctx.HTML(200, "dashboard/index")
}
