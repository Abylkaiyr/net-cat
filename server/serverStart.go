package server

import (
	"fmt"
	"log"
	"net"
)

const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "

func StartServer(hostname string, port int) {
	s := NewServer()
	go s.Run()

	addr := fmt.Sprintf("%s:%d", hostname, port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("connection error to the server with adress %s, define another adress\n", addr)
	}

	defer listener.Close()

	log.Printf("Listening the connections on the server %s", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error in accepting connection from the client: %s", err)
		} else {
			go s.handleRequest(conn)
		}
	}
}
