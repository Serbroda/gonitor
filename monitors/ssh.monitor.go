package monitors

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
)

type SSHMonitor struct {
	Host         string
	Port         string
	User         string
	Password     string
	Command      string
	ResultParser func(out string) bool
}

func (m *SSHMonitor) Monitor() bool {
	conf := &ssh.ClientConfig{
		User: m.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(m.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%s", m.Host, m.Port)
	client, err := ssh.Dial("tcp", addr, conf)
	if err != nil {
		fmt.Println("Failed to dial SSH")
		panic(err)
	}
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create SSH session")
		panic(err)
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(m.Command)
	if err != nil {
		fmt.Println("Failed to run command over SSH")
		panic(err)
	}
	out := b.String()
	fmt.Printf("%s:\n %s", m.Command, out)
	return m.ResultParser(out)
}
