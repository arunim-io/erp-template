package erp

import (
	"fmt"
	"net"
)

type Settings struct {
	Server struct {
		Host string
		Port int
	}
}

func (s Settings) GetServerAddress() string {
	return net.JoinHostPort(s.Server.Host, fmt.Sprint(s.Server.Port))
}

func GetSettings() Settings {
	s := Settings{}

	s.Server.Host = "0.0.0.0"
	s.Server.Port = 8000

	return s
}
