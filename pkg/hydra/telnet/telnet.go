package telnet

import (
	"github.com/w3security/goscan/pkg/kscan/lib/gotelnet"
)

func Check(addr, username, password string, port, serverType int) (bool, error) {
	client := gotelnet.New(addr, port)
	err := client.Connect()
	if err != nil {
		return false, err
	}
	defer client.Close()
	client.UserName = username
	client.Password = password
	client.ServerType = serverType
	err = client.Login()
	if err != nil {
		return false, err
	}
	return true, err
}
