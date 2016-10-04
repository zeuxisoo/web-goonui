package forms

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/binding"
)

type DashboardResultForm struct {
    Command     string      `form:"command" binding:"Required"`
    TargetIp    string      `form:"target_ip" binding:"Required"`
    Servers     []string    `form:"servers" binding:"Required"`
}

func (form *DashboardResultForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
    return validate(errs, ctx.Data, form)
}
