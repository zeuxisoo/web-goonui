package routes

import (
    "fmt"

    "github.com/zeuxisoo/go-goonui/app/kernels/context"
    "github.com/zeuxisoo/go-goonui/app/forms"
    "github.com/zeuxisoo/go-goonui/app/models"
)

func ServerCreate(ctx *context.Context) {
    ctx.HTML(200, "server/create")
}

func ServerStore(ctx *context.Context, form forms.CreateServerForm) {
    fmt.Printf("Name=%s, Host=%s, Port=%s, Username=%s, Password=%s, AuthMethod=%s\n",
        form.Name, form.Host, form.Port, form.Username, form.Password, form.AuthMethod)

    if ctx.HasError() {
        ctx.HTML(200, "server/create")
        return
    }

    err := models.CreateServer(&models.Server{
        Name: form.Name,
        Host: form.Host,
        Port: form.Port,
        Username: form.Username,
        Password: form.Password,
        AuthMethod: form.AuthMethod,
    })

    if err != nil {
        ctx.HTMLError(200, err.Error(), "server/create")
    }else{
        ctx.Flash.SuccessMsg = "Server Created!"
        ctx.Data["Flash"]    = ctx.Flash

        ctx.HTML(200, "server/create")
    }
}
