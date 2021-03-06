package main

import (
    "os"

    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-goonui/app/console"
)

const APP_VER = "0.1.0"

func main() {
    app := cli.NewApp()
    app.Name = "Goonui"
    app.Usage = "Goon user interface"
    app.Version = APP_VER
    app.Commands = []cli.Command{
        console.CmdWeb,
        console.CmdInstall,
    }
    app.Flags = append(app.Flags, []cli.Flag{}...)
    app.Run(os.Args)
}
