package main

import (
	"flag"
	"log"
	"net"

	"github.com/stinkyfingers/servers/tcp/server/handlers"
)

var (
	port = flag.String("port", "8888", "port")
)

func main() {
	flag.Parse()
	server, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlers.ConnectionHandler(connection)
	}
}
