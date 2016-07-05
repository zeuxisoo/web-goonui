
package cmd

import (
    "fmt"
    "net/http"

    "github.com/codegangsta/cli"

    "gopkg.in/macaron.v1"

    "github.com/zeuxisoo/go-goonui/modules/setting"
    "github.com/zeuxisoo/go-goonui/modules/log"
)

var CmdWeb = cli.Command{
    Name       : "serve",
    Usage      : "Start goon ui server",
    Description: `Goon ui is the only thing you need to run`,
    Action     : runWeb,
    Flags      : []cli.Flag{
        stringFlag("address, a", "127.0.0.1", "Custom address for server"),
        intFlag("port, p", 8080, "Custom port for server"),
    },
}

func runWeb(ctx *cli.Context) error {
    setting.NewSetting()

    //
    if ctx.IsSet("address") {
        setting.Address = ctx.String("address")
    }

    if ctx.IsSet("port") {
        setting.Port = ctx.Int("port")
    }

    //
    m := macaron.New()
    m.Get("/", func() string {
        return "Hello world!"
    })

    //
    addr := fmt.Sprintf("%s:%d", setting.Address, setting.Port)

    log.Infof("Listen: http://%s", addr)

    if err := http.ListenAndServe(addr, m); err != nil {
        log.Fatalf("Fail to start server: %v", err)
    }

    return nil
}
