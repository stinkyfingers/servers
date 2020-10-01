package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stinkyfingers/lambdify"
	"golang.org/x/net/websocket"
)

func main() {
	// http.ListenAndServe(":7000", mux())
	lambdaFunction := lambdify.Lambdify(mux())
	lambda.Start(lambdaFunction)
}

func mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", Status)
	mux.Handle("/ws", websocket.Handler(WS))
	return mux
}

// handlers
func WS(ws *websocket.Conn) {
	var msg string
	for {
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("recv err: ", err)
		}
		err = websocket.Message.Send(ws, "message received: "+msg)
		if err != nil {
			fmt.Println("send err: ", err)
		}
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status up"))
}

func Cors(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			return
		}
		fn(w, r)
	}
}
