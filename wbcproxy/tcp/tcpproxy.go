package tcp

import (
	"net"
	"strings"
	"encoding/hex"
	"runtime"
	"os"
	"fmt"
	"time"
)


func printableAddr(a net.Addr) string {
	return strings.Replace(a.String(), ":", "-", -1)
}

type Channel struct {
	from, to             net.Conn
	logger, binaryLogger chan []byte
	ack                  chan bool
}

func passThrough(c *Channel) {
	fromPeer := printableAddr(c.from.LocalAddr())
	toPeer := printableAddr(c.to.LocalAddr())

	b := make([]byte, 10240)
	offset := 0
	packet := 0
	for {
		n, err := c.from.Read(b)
		if err != nil {
			c.logger <- []byte(fmt.Sprintf("Disconnected from %s\n", fromPeer))
			break
		}
		if n > 0 {
			c.logger <- []byte(fmt.Sprintf("Received (#%d, %08X) %d bytes from %s\n",
				packet, offset, n, fromPeer))
			c.logger <- []byte(hex.Dump(b[:n]))
			c.binaryLogger <- b[:n]
			c.to.Write(b[:n])
			c.logger <- []byte(fmt.Sprintf("Sent (#%d) to %s\n", packet, toPeer))
			offset += n
			packet += 1
		}
	}
	c.from.Close()
	c.to.Close()
	c.ack <- true
}

func processConnection(local net.Conn, conn_n int, target string) {
	remote, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("Unable to connect to %s, %v\n", target, err)
	}

	localInfo := printableAddr(remote.LocalAddr())
	remoteInfo := printableAddr(remote.RemoteAddr())

	started := time.Now()

	logger := make(chan []byte)
	fromLogger := make(chan []byte)
	toLogger := make(chan []byte)

	ack := make(chan bool)

	go connectionLogger(logger, conn_n, localInfo, remoteInfo)
	go binaryLogger(fromLogger, conn_n, localInfo)
	go binaryLogger(toLogger, conn_n, remoteInfo)

	logger <- []byte(fmt.Sprintf("Connected to %s at %s\n", target,
		formatTime(started)))

	go passThrough(&Channel{remote, local, logger, toLogger, ack})
	go passThrough(&Channel{local, remote, logger, fromLogger, ack})

	<-ack
	<-ack

	finished := time.Now()
	duration := finished.Sub(started)
	logger <- []byte(fmt.Sprintf("Finished at %s, duration %s\n",
		formatTime(started), duration.String()))

	logger <- []byte{}
	fromLogger <- []byte{}
	toLogger <- []byte{}
}

func runTcpProxy(host, port, listenPort string) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	target := net.JoinHostPort(host, port)
	fmt.Printf("Start listening on port %s and forwarding data to %s\n",
		listenPort, target)

	ln, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		fmt.Printf("Unable to start listener, %v\n", err)
		os.Exit(1)
	}
	connN := 1
	for {
		if conn, err := ln.Accept(); err == nil {
			go processConnection(conn, connN, target)
			connN += 1
		} else {
			fmt.Printf("Accept failed, %v\n", err)
		}
	}
}