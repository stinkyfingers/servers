package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		Handler:   mux,
		TLSConfig: &tls.Config{},
		Addr:      ":8888",
		// WriteTimeout: time.Nanosecond,
	}
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	log.Print(string(b))
	resp := Response{
		Message: string(b),
		Time:    time.Now(),
	}
	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// CREATE CERT & KEY
// openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650
// ln -sf server.rsa.key server.key
// ln -sf server.rsa.crt server.crt

// REQUEST
// curl https://localhost:8888/enroll --insecure
