package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		w1, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(os.Stdout, w1)
		w1.Close()
	}
}
