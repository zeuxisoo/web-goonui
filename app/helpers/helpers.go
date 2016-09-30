package helpers

import (
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func GetRandomString(length int) string {
    var words = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+_)(*&$#@!")

    runes := make([]rune, length)
    for i := range runes {
        runes[i] = words[rand.Intn(len(words))]
    }

    return string(runes)
}
