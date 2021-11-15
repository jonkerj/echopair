package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Echo struct {
	ListenAddr string
	Logger     *log.Logger
}

func New(port int, log *log.Logger) Echo {
	server := &Echo{
		ListenAddr: fmt.Sprintf(":%d", port),
		Logger:     log,
	}

	return *server
}

func (e *Echo) Serve() {
	listener, err := net.Listen("tcp", e.ListenAddr)
	e.Logger.Printf("Listening on %s", e.ListenAddr)
	for i := 0; err == nil; i++ {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Could not accept: %v", err)
			continue
		}
		go e.HandleClient(conn, i)
	}
}

func (e *Echo) HandleClient(conn net.Conn, i int) {
	defer conn.Close()

	e.Logger.Printf("%d: %v <-> %v: connect", i, conn.LocalAddr(), conn.RemoteAddr())
	b := bufio.NewReader(conn)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}
		conn.Write(line)
	}

	e.Logger.Printf("%d: close", i)
}
