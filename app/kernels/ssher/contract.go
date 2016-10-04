package ssher

import (
    "golang.org/x/crypto/ssh"
)

type Config struct {
    Host        string
    Port        int
    User        string
    Password    string
    PrivateKey  string
}

type Authenticator interface {
    SshClient() (*ssh.Client, error)
    SetConfig(*Config)
}
