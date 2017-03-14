package handlers

import (
	"io/ioutil"
	"log"
	"net"
)

const BUFFER_SIZE = 1024

func ConnectionHandler(conn net.Conn) {
	b, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(b))

}
