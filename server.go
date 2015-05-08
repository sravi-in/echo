package main

import "net"

func acceptConnection(c net.Conn) {
	var buf []byte
	_, err := c.Read(buf)
	if err != nil {
		panic(err)
	}
	c.Write(buf)
}

func main() {
	ln, err := net.Listen("tcp", ":2345")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go acceptConnection(c)
	}
}
