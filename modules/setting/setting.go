package setting

import (
    "gopkg.in/ini.v1"

    "github.com/zeuxisoo/go-goonui/modules/log"
)

var (
    Address     string
    Port        string

    Configure   *ini.File
)

func NewSetting() {
    Configure, err := ini.Load("conf/app.ini")

    if err != nil {
        log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
    }

    sectionServer := Configure.Section("server")

    Address = sectionServer.Key("ADDRESS").MustString("127.0.0.1")
    Port    = sectionServer.Key("PORT").MustString("8080")
}
