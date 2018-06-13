package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var port = flag.String("port", "9876", "listen port")

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Print("running on port ", *port)

	log.Fatal(http.ListenAndServe(":"+*port, mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nreceived from: %s; URI: %s\n", r.RemoteAddr, r.URL)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(string(b))

	w.Write([]byte("OK"))
}
