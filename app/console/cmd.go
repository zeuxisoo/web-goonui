package console

import (
    "github.com/codegangsta/cli"
)

func stringFlag(name string, value string, usage string) cli.StringFlag {
    return cli.StringFlag{
        Name : name,
        Value: value,
        Usage: usage,
    }
}

func intFlag(name string, value int, usage string) cli.IntFlag {
    return cli.IntFlag{
        Name :  name,
        Value: value,
        Usage: usage,
    }
}
