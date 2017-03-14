package main

import (
	"flag"
	"log"
	"net"
	"os"
)

var (
	port = flag.String("port", "8888", "port")
	ip   = flag.String("ip", "localhost", "ip")
)

func main() {
	conn, err := net.Dial("tcp", *ip+":"+*port)
	if err != nil {
		log.Fatal("CONN ERR ", err)
	}

	msg := os.Args[1]

	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}
