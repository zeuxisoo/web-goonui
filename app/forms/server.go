package forms

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/binding"
)

type CreateServerForm struct {
    Name       string `form:"name" binding:"Required;MaxSize(120)"`
    Host       string `form:"host" binding:"Required;MaxSize(120)"`
    Port       string `form:"port" binding:"Required;MaxSize(10)"`
    Username   string `form:"username" binding:"Required;MaxSize(50)"`
    Password   string `form:"password" binding:"Required"`
    AuthMethod string `form:"auth_method" binding:"Required`
}

func (form *CreateServerForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
    return validate(errs, ctx.Data, form)
}
