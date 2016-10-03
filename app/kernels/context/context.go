package context

import (
    "fmt"
    "strings"
    "html/template"

    "gopkg.in/macaron.v1"
    "github.com/go-macaron/csrf"
    "github.com/go-macaron/session"

    "github.com/zeuxisoo/go-goonui/app/kernels/setting"
    "github.com/zeuxisoo/go-goonui/app/models"
)

type Context struct {
    *macaron.Context

    csrf        csrf.CSRF
    Flash       *session.Flash
    Session     session.Store

    IsSigned    bool

    User        models.User
}

func (c *Context) HasError() bool {
    hasErr, ok := c.Data["HasError"]

    if !ok {
        return false
    }

    c.Flash.ErrorMsg = c.Data["ErrorMessage"].(string)
    c.Data["Flash"]  = c.Flash

    return hasErr.(bool)
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

func (c *Context) HTMLError(status int, message string, name string) {
    c.Flash.ErrorMsg = message
    c.Data["Flash"]  = c.Flash

    c.Context.HTML(status, name)
}

func (c *Context) HTML(status int, name string) {
    c.Context.HTML(status, name)
}

func SignInUserBySession(ctx *macaron.Context, sess session.Store) (models.User, error) {
    userId := sess.Get("user_id")

    if userId == nil {
        userId = int64(0)
    }

    user, err := models.FindUserById(userId.(int64))

    if err != nil {
        return user, err
    }else{
        return user, nil
    }
}

func Contexter() macaron.Handler {
    return func(c *macaron.Context, s session.Store, f *session.Flash, x csrf.CSRF) {
        ctx := &Context{
            Context: c,
            csrf   : x,
            Flash  : f,
            Session: s,
        }

        ctx.User, _ = SignInUserBySession(ctx.Context, ctx.Session)

        if ctx.User.Username != "" {
            ctx.IsSigned = true
        }else{
            ctx.IsSigned = false
        }

        ctx.Data["Link"]          = setting.AppUrl + strings.TrimSuffix(ctx.Req.URL.Path, "/")
        ctx.Data["CsrfToken"]     = x.GetToken()
        ctx.Data["CsrfTokenHtml"] = template.HTML(`<input type="hidden" name="_csrf" value="` + x.GetToken() + `">`)

        c.Map(ctx)
    }
}
