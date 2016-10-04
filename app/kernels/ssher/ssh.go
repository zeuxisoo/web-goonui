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

func (s *Ssh) RunCommand(command string) (result string) {
    sshClient, _  := s.authenticator.SshClient()
    sshSession, _ := s.createSshSession(sshClient)

    buffer, _ := sshSession.CombinedOutput(command)

    sshSession.Close()
    sshClient.Close()

    return string(buffer)
}

func (s *Ssh) createSshSession(sshClient *ssh.Client) (*ssh.Session, error) {
    session, err := sshClient.NewSession()

    if err != nil {
        return nil, errors.New("Failed to create session")
    }

    return session, nil
}
