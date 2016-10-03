package routes

import (
    "github.com/jinzhu/gorm"

    "github.com/zeuxisoo/go-goonui/app/kernels/context"
    "github.com/zeuxisoo/go-goonui/app/kernels/setting"
    "github.com/zeuxisoo/go-goonui/app/forms"
    "github.com/zeuxisoo/go-goonui/app/models"
)

func Home(ctx *context.Context) {
    ctx.HTML(200, "index")
}

func DoSignIn(ctx *context.Context, form forms.SignInForm) {
    if ctx.HasError() {
        ctx.HTML(200, "index")
        return
    }

    user, err := models.SignInUser(form.Username, form.Password)

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            ctx.HTMLError(200, "Not found user", "index")
        }else if err.Error() == "PASSWORD_NOT_MATCH" {
            ctx.HTMLError(200, "Password not match", "index")
        }else{
            ctx.Handle(500, "DoSignIn", err)
        }
        return
    }

    ctx.Session.Set("user_id",  user.ID)
    ctx.Session.Set("username", user.Username)

    ctx.Redirect(setting.AppSubUrl + "/dashboard")
}
