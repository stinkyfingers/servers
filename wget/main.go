package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	file = flag.String("f", "", "file location")
)

func main() {
	flag.Parse()
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Print("Running ", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(*file)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
