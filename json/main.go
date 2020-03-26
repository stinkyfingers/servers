package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := struct {
			Name string `json:"name"`
			Age string `json:"age"`
		}{
			"Jim",
			"55",
		}
		j, err := json.Marshal(t)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(j)
	})
	log.Fatal(http.ListenAndServe(":5000", mux))
}
