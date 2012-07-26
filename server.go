package kartoffel

import (
	"fmt"
	"io"
	"net"
)

type ServerState int

const (
	Accepting ServerState = iota
	ShuttingDown
)

type ServerStatus struct {
	State ServerState
	net.Addr
}

// Main entry point for server
func Listen(port int, status chan<- ServerStatus) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	status <- ServerStatus{Accepting, l.Addr()}
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)

	}

	return nil
}
