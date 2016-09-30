package context

import (
    "fmt"
    "strings"
    "html/template"

    "gopkg.in/macaron.v1"
    "github.com/go-macaron/csrf"
    "github.com/go-macaron/session"

    "github.com/zeuxisoo/go-goonui/app/kernels/setting"
)

type Context struct {
    *macaron.Context

    csrf    csrf.CSRF
    Flash   *session.Flash
    Session session.Store
}

func (c *Context) Handle(status int, title string, err error) {
    if err != nil {
        if macaron.Env != macaron.PROD {
            c.Data["ErrorMessage"] = err
        }
    }

    switch status {
        case 404:
            c.Data["Title"] = "Page Not Found"
        case 500:
            c.Data["Title"] = "Internal Server Error"
    }

    c.HTML(status, fmt.Sprintf("status/%d", status))
}

func (c *Context) HTML(status int, name string) {
    c.Context.HTML(status, name)
}

func Contexter() macaron.Handler {
    return func(c *macaron.Context, s session.Store, f *session.Flash, x csrf.CSRF) {
        ctx := &Context{
            Context: c,
            csrf   : x,
            Flash  : f,
            Session: s,
        }

        ctx.Data["Link"]          = setting.AppUrl + strings.TrimSuffix(ctx.Req.URL.Path, "/")
        ctx.Data["CsrfToken"]     = x.GetToken()
        ctx.Data["CsrfTokenHtml"] = template.HTML(`<input type="hidden" name="_csrf" value="` + x.GetToken() + `">`)

        c.Map(ctx)
    }
}
