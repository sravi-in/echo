package main

import (
	"io"
	"log"
	"net"
)

const (
	listenPort   = "2345"
	maxEchoBytes = 1024 * 1024 // 1 MB
)

func acceptConnection(c net.Conn) {
	log.Println("Accepted connection from", c.RemoteAddr())
	n, err := io.CopyN(c, c, maxEchoBytes)
	if err != nil && err != io.EOF {
		log.Print("Early termination:", err)
	}
	log.Println("Terminating connection from", c.RemoteAddr(),
		"after echoing", n, "bytes")
	c.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	log.Print("Echo server listening on", ln.Addr())
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal("Failed to accept: ", err)
		}
		go acceptConnection(c)
	}
}
