package models

import (
    "time"
)

type Server struct {
    ID          int64       `gorm:"primary_key;AUTO_INCREMENT"`
    Name        string      `gorm:"type:varchar(120)"`
    Host        string      `gorm:"type:varchar(120)"`
    Port        string      `gorm:"type:varchar(10)"`
    Username    string      `gorm:"type:varchar(50)"`
    Password    string      `gorm:"type:varchar(255)"`
    AuthMethod  string      `gorm:"type:varchar(30)"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func CreateServer(server *Server) error {
    if err := db.Create(&server).Error; err != nil {
        return err
    }

    return nil
}
