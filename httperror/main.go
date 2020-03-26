package main

import (
	"flag"
	"log"
	"net/http"
)

var code = flag.Int("code", 403, "status code")

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "this is an error", *code)
		return
	})
	log.Fatal(http.ListenAndServe(":9000", nil))
}
