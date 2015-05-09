package main

import (
	"net"
	"fmt"
	"io"
	"log"
)

func acceptConnection(c net.Conn) {
	fmt.Println("Accepted connection", c)
	io.Copy(c, c)
	fmt.Println("Terminating connection", c)
	c.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":2345")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal("Failed to accept: ", err)
		}
		go acceptConnection(c)
	}
}
