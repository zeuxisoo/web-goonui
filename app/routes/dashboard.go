package routes

import (
    "strconv"

    "github.com/zeuxisoo/go-goonui/app/kernels/context"
    "github.com/zeuxisoo/go-goonui/app/kernels/ssher"
    "github.com/zeuxisoo/go-goonui/app/models"
    "github.com/zeuxisoo/go-goonui/app/forms"
)

type serverResult struct {
    server   string
    response string
}

func Dashboard(ctx *context.Context) {
    servers, err := models.FindAllServers()

    if err != nil {
        ctx.HTMLError(200, err.Error(), "dashboard/index")
    }else{
        ctx.Data["Servers"] = servers

        ctx.HTML(200, "dashboard/index")
    }
}

func DashboardResult(ctx *context.Context, form forms.DashboardResultForm) {
    servers, err := models.FindAllServersByIds(form.Servers)

    if err != nil {
        ctx.HTMLError(200, err.Error(), "dashboard/index")
    }else{
        var results []serverResult

        for _, server := range servers {
            port, _ := strconv.Atoi(server.Port)

            authenticator := new(ssher.PasswordAuthenticator)
            authenticator.SetConfig(&ssher.Config{
                Host      : server.Host,
                Port      : port,
                User      : server.Username,
                Password  : server.Password,
                PrivateKey: "",
            })

            sshAgent := ssher.NewSsh()
            sshAgent.SetAuthenticator(authenticator)

            result := serverResult{
                server  : server.Name,
                response: sshAgent.RunCommand("host www.yahoo.com.hk"),
            }

            results = append(results, result)
        }

        ctx.Data["Command"] = form.Command
        ctx.Data["Results"] = results

        ctx.HTML(200, "dashboard/result")
    }
}
