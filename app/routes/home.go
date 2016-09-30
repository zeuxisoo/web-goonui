package routes

import (
    "github.com/zeuxisoo/go-goonui/app/kernels/context"
    "github.com/zeuxisoo/go-goonui/app/kernels/setting"
)

func Home(ctx *context.Context) {
    ctx.HTML(200, "index")
}

func DoSignIn(ctx *context.Context) {
    ctx.Redirect(setting.AppSubUrl + "/")
}
