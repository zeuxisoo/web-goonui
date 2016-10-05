package console

import (
    "fmt"
    "path"
    "net/http"

    "github.com/codegangsta/cli"

    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "github.com/go-macaron/csrf"
    "github.com/go-macaron/pongo2"
    "github.com/go-macaron/binding"

    "github.com/zeuxisoo/go-goonui/app/kernels/setting"
    "github.com/zeuxisoo/go-goonui/app/kernels/log"
    "github.com/zeuxisoo/go-goonui/app/kernels/context"
    "github.com/zeuxisoo/go-goonui/app/kernels/middleware"
    "github.com/zeuxisoo/go-goonui/app/models"
    "github.com/zeuxisoo/go-goonui/app/routes"
    "github.com/zeuxisoo/go-goonui/app/forms"
)

var CmdWeb = cli.Command{
    Name       : "web",
    Usage      : "Start goon ui server",
    Description: `Goon ui is the only thing you need to run`,
    Action     : runWeb,
    Flags      : []cli.Flag{
        stringFlag("address, a", "127.0.0.1", "Custom address for server"),
        stringFlag("port, p", "8080", "Custom port for server"),
    },
}

func runWeb(ctx *cli.Context) error {
    setting.NewSetting()
    setting.NewSessionService()

    //
    models.LoadConfigs()

    if err := models.NewDB(); err != nil {
        log.Fatalf("Failed to initialize ORM engine: %v", err)
    }

    //
    if ctx.IsSet("address") {
        setting.Address = ctx.String("address")
    }

    if ctx.IsSet("port") {
        setting.Port = ctx.String("port")
    }

    //
    m := macaron.New()
    m.Use(macaron.Logger())
    m.Use(macaron.Recovery())

    m.Use(macaron.Static(
        path.Join(setting.StaticRootPath, "public"),
        macaron.StaticOptions{
            SkipLogging: false,
        },
    ))

    m.Use(pongo2.Pongoer(pongo2.Options{
        Directory      : path.Join(setting.StaticRootPath, "app/templates"),
        Extensions     : []string{ ".tmpl", ".html" },
        Charset        : "UTF-8",
        IndentJSON     : true,
        IndentXML      : true,
        HTMLContentType: "text/html",
    }))

    m.Use(session.Sessioner(setting.SessionConfig))

    m.Use(csrf.Csrfer(csrf.Options{
        Secret    : setting.SecretKey,
        Cookie    : setting.CsrfCookieName,
        SetCookie : true,
        Header    : "X-Csrf-Token",
        CookiePath: setting.AppSubUrl,
    }))

    m.Use(context.Contexter())

    //
    bindIgnErr := binding.BindIgnErr

    //
    m.Get("/", routes.Home)
    m.Post("/signin", csrf.Validate, bindIgnErr(forms.SignInForm{}), routes.DoSignIn)

    m.Get("/dashboard", middleware.RequreSignIn, routes.Dashboard)
    m.Post("/dashboard/result", middleware.RequreSignIn, csrf.Validate, bindIgnErr(forms.DashboardResultForm{}), routes.DashboardResult)

    m.Group("/server", func() {
        m.Get("/", routes.ServerIndex)
        m.Get("/create", routes.ServerCreate)
        m.Post("/store", csrf.Validate, bindIgnErr(forms.CreateServerForm{}), routes.ServerStore)
        m.Get("/edit/:serverid", routes.ServerEdit)
        m.Post("/update/:serverid", csrf.Validate, bindIgnErr(forms.EditServerForm{}), routes.ServerUpdate)
    }, middleware.RequreSignIn)

    //
    addr := fmt.Sprintf("%s:%s", setting.Address, setting.Port)

    log.Infof("Listen   : http://%s", addr)
    log.Infof("AppPath  : %s", setting.AppPath)
    log.Infof("SecretKey: %s", setting.SecretKey)

    if err := http.ListenAndServe(addr, m); err != nil {
        log.Fatalf("Fail to start server: %v", err)
    }

    return nil
}
