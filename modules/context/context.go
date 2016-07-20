package context

import (
    "strings"

    "gopkg.in/macaron.v1"

    "github.com/zeuxisoo/go-goonui/modules/setting"
)

type Context struct {
    *macaron.Context
}

func (c *Context) HTML(status int, name string) {
    c.Context.HTML(status, name)
}

func Contexter() macaron.Handler {
    return func(c *macaron.Context) {
        ctx := &Context{
            Context: c,
        }

        ctx.Data["Link"] = setting.AppUrl + strings.TrimSuffix(ctx.Req.URL.Path, "/")

        c.Map(ctx)
    }
}
