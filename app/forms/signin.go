package forms

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/binding"
)

type SignInForm struct {
    Username string `form:"username" binding:"Required;MaxSize(30)"`
    Password string `form:"password" binding:"Required;MaxSize(120)"`
}

func (form *SignInForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
    return validate(errs, ctx.Data, form)
}
