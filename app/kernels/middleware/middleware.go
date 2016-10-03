package middleware

import (
    "github.com/zeuxisoo/go-goonui/app/kernels/context"
    "github.com/zeuxisoo/go-goonui/app/kernels/setting"
)

func RequreSignIn(ctx *context.Context) {
    if ctx.IsSigned == false {
        ctx.Redirect(setting.AppSubUrl + "/")
    }
}
