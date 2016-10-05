package routes

import (
    "github.com/zeuxisoo/go-goonui/app/kernels/context"
    "github.com/zeuxisoo/go-goonui/app/forms"
    "github.com/zeuxisoo/go-goonui/app/models"
)

func ServerIndex(ctx *context.Context) {
    servers, err := models.FindAllServers()

    if err != nil {
        ctx.HTMLError(200, err.Error(), "server/index")
    }else{
        ctx.Data["Servers"] = servers

        ctx.HTML(200, "server/index")
    }
}

func ServerCreate(ctx *context.Context) {
    ctx.HTML(200, "server/create")
}

func ServerStore(ctx *context.Context, form forms.CreateServerForm) {
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

func ServerEdit(ctx *context.Context) {
    server, err := models.FindServerById(ctx.ParamsInt64(":serverid"))

    if err != nil {
        ctx.HTMLError(200, err.Error(), "server/create")
    }else if server.Name == "" {
        ctx.HTMLError(200, "The server information incorrect", "server/create")
    }else{
        ctx.Data["Server"] = server

        ctx.HTML(200, "server/edit")
    }
}

func ServerUpdate(ctx *context.Context, form forms.EditServerForm) {
    server, err := models.FindServerById(ctx.ParamsInt64(":serverid"))

    if err != nil {
        ctx.HTMLError(200, err.Error(), "server/create")
    }else if server.Name == "" {
        ctx.HTMLError(200, "The server information incorrect", "server/create")
    }else if ctx.HasError() {
        ctx.Data["Server"] = server
        ctx.HTML(200, "server/edit")
    }else{
        server.Name       = form.Name
        server.Host       = form.Host
        server.Port       = form.Port
        server.Username   = form.Username
        server.Password   = form.Password
        server.AuthMethod = form.AuthMethod

        err := models.UpdateServerById(&server)

        if err != nil {
            ctx.Flash.ErrorMsg = "Cannot update server!"
            ctx.Data["Flash"]  = ctx.Flash
        }else{
            ctx.Flash.SuccessMsg = "Server Updated!"
            ctx.Data["Flash"]    = ctx.Flash
        }

        ctx.Data["Server"] = server
        ctx.HTML(200, "server/edit")
    }
}
