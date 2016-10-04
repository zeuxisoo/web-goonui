package ssher

import (
    "time"
    "strconv"
    "errors"

    "golang.org/x/crypto/ssh"
)

type PasswordAuthenticator struct {
    config *Config
}

// Implement
func (p *PasswordAuthenticator) SshClient() (*ssh.Client, error) {
    clientConfig := &ssh.ClientConfig{
        User: p.config.User,
        Auth: []ssh.AuthMethod{
            ssh.Password(p.config.Password),
        },
        Timeout: 30 * time.Second,
    }

    address := p.config.Host + ":" + strconv.Itoa(p.config.Port)

    client, err := ssh.Dial("tcp", address, clientConfig)

    if err != nil {
        return nil, errors.New("Failed to create client")
    }

    return client, nil
}

func (p *PasswordAuthenticator) SetConfig(config *Config) {
    p.config = config
}
