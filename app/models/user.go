package models

import (
    "fmt"
    "time"
    "errors"
    "strings"

    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"
)

type User struct {
    ID          int64       `gorm:"primary_key;AUTO_INCREMENT"`
    Username    string      `gorm:"type:varchar(50);unique_index"`
    Password    string      `gorm:"type:varchar(100)"`
    Email       string      `gorm:"type:varchar(120);unique_index"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func (user *User) EncryptPassword() error {
    hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)

    if err != nil {
        return fmt.Errorf("Cannot encrypt user password: %v", err)
    }else{
        user.Password = string(hashBytes)

        return nil
    }
}

//
func IsUsernameExist(username string) bool {
    err := db.Where("username = ?", username).First(&User{}).Error

    return err != gorm.ErrRecordNotFound
}

func IsEmailExist(email string) bool {
    err := db.Where("email = ?", email).First(&User{}).Error;

    return err != gorm.ErrRecordNotFound
}

func IsEqualsPassword(hashedPassword string, plainPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

    return err == nil
}

//
func CreateUser(user *User) error {
    if len(user.Username) == 0 {
        return errors.New("Username cannot blank")
    }else if len(user.Email) == 0 {
        return errors.New("Email cannot blank")
    }else if len(user.Password) < 8 {
        return errors.New("Password length must more than 8 char")
    }else if IsUsernameExist(user.Username) {
        return errors.New("Username already exists")
    }else if IsEmailExist(user.Email) {
        return errors.New("Email already exists")
    }else{
        user.EncryptPassword()

        if err := db.Create(&user).Error; err != nil {
            return err
        }

        return nil
    }
}

func SignInUser(username string, password string) (User, error) {
    var user User
    var whereCondition User

    if strings.Contains(username, "@") {
        whereCondition = User{ Email: username }
    }else{
        whereCondition = User{ Username: username }
    }

    orm := db.First(&user, whereCondition)

    if err := orm.Error; err != nil {
        return user, err
    }else if IsEqualsPassword(user.Password, password) == false {
        return user, errors.New("PASSWORD_NOT_MATCH")
    }else{
        return user, nil
    }
}
