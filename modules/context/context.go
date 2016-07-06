package context

import (
    "gopkg.in/macaron.v1"
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

        c.Map(ctx)
    }
}
