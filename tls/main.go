package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	NodeKey string `json:"node_key"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/enroll", enrollHandler)
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(":8888", "server.crt", "server.key", mux))
}

func enrollHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	log.Print(string(b))
	resp := Response{"test"}
	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	log.Print(string(b))
	l := struct{}{}
	j, _ := json.Marshal(l)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// CREATE CERT & KEY
// openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650
// ln -sf server.rsa.key server.key
// ln -sf server.rsa.crt server.crt
