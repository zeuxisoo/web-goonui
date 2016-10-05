package routes

import (
    "fmt"
    "strconv"
    "strings"

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
        var command string

        switch strings.ToLower(form.Command) {
            case "ping":
                command = fmt.Sprintf("ping -c 4 -t 15 %s", form.TargetIp)
            case "host":
                command = fmt.Sprintf("host %s", form.TargetIp)
            case "traceroute":
                command = fmt.Sprintf("traceroute -n -m 30 %s", form.TargetIp)
            case "mtr":
                command = fmt.Sprintf("mtr -rw %s", form.TargetIp)
            case "nslookup":
                command = fmt.Sprintf("nslookup %s", form.TargetIp)
            default:
                command = fmt.Sprintf("Cannot find related command: %s", form.TargetIp)
        }

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

            response, err := sshAgent.RunCommand(command)

            if err != nil {
                response = err.Error()
            }

            result := serverResult{
                server  : server.Name,
                response: response,
            }

            results = append(results, result)
        }

        ctx.Data["Command"] = command
        ctx.Data["Results"] = results

        ctx.HTML(200, "dashboard/result")
    }
}
