package routes

import (
    "github.com/zeuxisoo/go-goonui/modules/context"
    "github.com/zeuxisoo/go-goonui/modules/setting"
)

func Home(ctx *context.Context) {
    ctx.HTML(200, "index")
}

func DoSignIn(ctx *context.Context) {
    ctx.Redirect(setting.AppSubUrl + "/")
}
