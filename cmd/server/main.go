package main

import (
	"fmt"
	"netcat/server"
	"os"
	"strconv"
)

func main() {
	host := "0.0.0.0"
	port := 8989

	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	if len(os.Args) == 2 {
		var err error

		port, err = strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println("Invalid port number, Please try format 8989")
			return
		}
	}

	server.StartServer(host, port)
}
