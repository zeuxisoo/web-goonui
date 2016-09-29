package models

import (
    "os"
    "fmt"
    "path"

    "github.com/jinzhu/gorm"

    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/zeuxisoo/go-goonui/modules/setting"
    "github.com/zeuxisoo/go-goonui/models/migrations"
)

var (
    db *gorm.DB

    DBConfigure struct {
        Driver      string
        Host        string
        Username    string
        Password    string
        Name        string
        SSLMode     string
        LogMode     bool
    }
)

func LoadConfigs() {
    section := setting.Configure.Section("database")

    DBConfigure.Driver   = section.Key("DRIVER").String()
    DBConfigure.Host     = section.Key("HOST").String()
    DBConfigure.Username = section.Key("USERNAME").String()
    DBConfigure.Password = section.Key("PASSWORD").String()
    DBConfigure.Name     = section.Key("NAME").String()
    DBConfigure.SSLMode  = section.Key("SSLMODE").String()
    DBConfigure.LogMode  = section.Key("LOG_MODE").MustBool(false)
}

func CreateDatabase() {
    var tables []interface{}

    tables = append(
        tables,
        new(User),
    )

    for _, table := range tables {
        if err := db.AutoMigrate(table).Error; err != nil {
            fmt.Errorf("Failed to create table schema: %v", err)
        }
    }
}

func getDB() (*gorm.DB, error) {
    dsn := ""

    switch(DBConfigure.Driver) {
        case "mysql":
            dsn = fmt.Sprintf(
                "%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
                DBConfigure.Username,
                DBConfigure.Password,
                DBConfigure.Host,
                DBConfigure.Name,
            )
        case "postgres":
            dsn = fmt.Sprintf(
                "host=%s user=%s dbname=%s sslmode=%s password=%s",
                DBConfigure.Host,
                DBConfigure.Username,
                DBConfigure.Name,
                DBConfigure.SSLMode,
                DBConfigure.Password,
            )
        case "sqlite3":
            if err := os.MkdirAll(path.Dir(DBConfigure.Host), os.ModePerm); err != nil {
                return nil, fmt.Errorf("Failed to create database directories: %v", err)
            }

            dsn = "file:" + DBConfigure.Host + "?cache=shared&mode=rwc"
        default:
            return nil, fmt.Errorf("Unknown database type: %s", DBConfigure.Driver)
    }

    return gorm.Open(DBConfigure.Driver, dsn)
}

func SetDB() (err error) {
    db, err = getDB()

    if err != nil {
        return fmt.Errorf("Cannot connect to database: %v", err)
    }

    db.LogMode(DBConfigure.LogMode)

    return nil
}

func NewDB() (err error) {
    if err = SetDB(); err != nil {
        return err
    }

    // db.AutoMigrate()
    if err = migrations.Migrate(db); err != nil {
        return fmt.Errorf("migrate: %v", err)
    }

    return nil
}
