package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	addr  = flag.String("a", "127.0.0.1:8888", "rhost")
	local = flag.String("l", "127.0.0.1:8080", "lhost")
)

func main() {
	flag.Parse()
	ServerAddr, err := net.ResolveUDPAddr("udp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer Conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
