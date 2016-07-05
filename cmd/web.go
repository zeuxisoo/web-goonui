
package cmd

import (
    "fmt"
    "path"
    "net/http"

    "github.com/codegangsta/cli"

    "gopkg.in/macaron.v1"
    "github.com/go-macaron/pongo2"

    "github.com/zeuxisoo/go-goonui/modules/setting"
    "github.com/zeuxisoo/go-goonui/modules/log"
    "github.com/zeuxisoo/go-goonui/routes"
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
        Directory      : path.Join(setting.StaticRootPath, "templates"),
        Extensions     : []string{ ".tmpl", ".html" },
        Charset        : "UTF-8",
        IndentJSON     : true,
        IndentXML      : true,
        HTMLContentType: "text/html",
    }))

    //
    m.Get("/", routes.Home)

    //
    addr := fmt.Sprintf("%s:%s", setting.Address, setting.Port)

    log.Infof("Listen: http://%s", addr)

    if err := http.ListenAndServe(addr, m); err != nil {
        log.Fatalf("Fail to start server: %v", err)
    }

    return nil
}
