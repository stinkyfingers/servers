package main

import (
	"flag"
	"log"
	"net"

	"github.com/stinkyfingers/servers/udp/serve/handlers"
)

var (
	port = flag.String("port", "8888", "port")
)

func main() {
	flag.Parse()
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	server, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	buf := make([]byte, 1024)

	for {
		_, addr, err := server.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}

		handlers.ConnectionHandler(buf, addr.String())
	}
}
