package handlers

import (
	"log"
)

func ConnectionHandler(b []byte, addr string) {

	log.Print("ADDR: ", addr, "  MSG: ", string(b))

}
