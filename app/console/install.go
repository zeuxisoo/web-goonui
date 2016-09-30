package console

import (
    "github.com/codegangsta/cli"

    "github.com/zeuxisoo/go-goonui/app/kernels/setting"
    "github.com/zeuxisoo/go-goonui/app/kernels/log"
    "github.com/zeuxisoo/go-goonui/app/models"
)

var CmdInstall = cli.Command{
    Name       : "install",
    Usage      : "Install goon ui data",
    Description: `Install goon ui program and database`,
    Action     : installGoon,
    Flags      : []cli.Flag{
    },
}

func installGoon(ctx *cli.Context) error {
    // Read configuration from file
    setting.NewSetting()

    // Load database configration
    models.LoadConfigs()

    if err := models.NewDB(); err != nil {
        log.Fatalf("Failed to initialize ORM engine: %v", err)
    }

    // Create all tables
    models.CreateDatabase()

    // Create default user
    user := &models.User{
        Username: "test",
        Password: "testtest",
        Email   : "test@test.com",
    }

    if err := models.CreateUser(user); err != nil {
        log.Fatalf("Failed to create default user account: %v", err)
    }

    return nil
}
