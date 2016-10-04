package ssher

import (
    "errors"

    "golang.org/x/crypto/ssh"
)

type Ssh struct {
    authenticator   Authenticator
}

func NewSsh() (s *Ssh) {
    return &Ssh{}
}

func (s *Ssh) SetAuthenticator(authenticator Authenticator) {
    s.authenticator = authenticator
}

func (s *Ssh) RunCommand(command string) (result string, err error) {
    if sshClient, err := s.authenticator.SshClient(); err != nil {
        return "", err
    }else{
        sshSession, err := s.createSshSession(sshClient);

        if err != nil {
            return "", err
        }

        buffer, _ := sshSession.CombinedOutput(command)

        sshSession.Close()
        sshClient.Close()

        return string(buffer), nil
    }
}

func (s *Ssh) createSshSession(sshClient *ssh.Client) (*ssh.Session, error) {
    session, err := sshClient.NewSession()

    if err != nil {
        return nil, errors.New("Failed to create session")
    }

    return session, nil
}
