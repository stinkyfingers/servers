package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

// NOTE: this does not work

var (
	url = "https://server.john-shenk.com/lambdawebsocket/ws"
	// url = "ws://localhost:7000/ws"
)

func main() {

	ws, err := websocket.Dial(url, "", "http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			text := scanner.Text()
			err = websocket.Message.Send(ws, text)
			if err != nil {
				log.Fatal(err)
			}

		}
	}()

	for {
		var msg string
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Print(err)
			break
		}
		fmt.Println("Returned message: ", msg)
	}

}
