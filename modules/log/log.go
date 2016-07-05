package log

import (
    "github.com/Sirupsen/logrus"
)

var (
    log = logrus.New()
)

func init() {
    log.Formatter = new(logrus.TextFormatter)
    log.Level     = logrus.DebugLevel
}

func Info(args ...interface{}) {
    log.Info(args...)
}

func Infof(format string, args ...interface{}) {
    log.Infof(format, args...)
}

func Fatalf(format string, args ...interface{}) {
    log.Fatalf(format, args...)
}
