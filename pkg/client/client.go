package client

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

type Echo struct {
	RemoteAddr *net.TCPAddr
	Logger     *log.Logger
	Interval   time.Duration
}

func New(address string, log *log.Logger, interval string) Echo {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Panicf("Could not resolve '%s': %v", address, err)
	}

	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Panicf("Could not parse duration '%s': %v", interval, err)
	}

	client := &Echo{
		RemoteAddr: tcpAddr,
		Logger:     log,
		Interval:   d,
	}

	return *client
}

func (e *Echo) EchoOnce() {
	e.Logger.Printf("Connecting to %v", e.RemoteAddr)
	conn, err := net.DialTCP("tcp", nil, e.RemoteAddr)
	if err != nil {
		e.Logger.Printf("Could not dial tcp: %v", err)
		return
	}

	defer conn.Close()

	t := fmt.Sprintf("test %d\n", rand.Intn(65536))

	_, err = conn.Write([]byte(t))
	if err != nil {
		e.Logger.Printf("Could not write: %v", err)
		return
	}

	reply := make([]byte, 1024)
	bytes, err := conn.Read(reply)
	if err != nil {
		e.Logger.Printf("Could not read: %v", err)
		return
	}

	r := string(reply)[:bytes]

	if r != t {
		e.Logger.Printf("Echo differs from test string. Got '%s' expected '%s'", r, t)
	} else {
		e.Logger.Printf("Perfect echo")
	}
}

func (e *Echo) EchoMany() {
	e.Logger.Printf("Invoking echo every %v", e.Interval)
	ticker := time.NewTicker(e.Interval)
	e.EchoOnce() // first tick is after the interval
	for range ticker.C {
		e.EchoOnce()
	}
}
