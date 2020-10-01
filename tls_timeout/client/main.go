package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://localhost:8888", nil)
	if err != nil {
		log.Fatal(err)
	}
	cli := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				// Timeout: time.Nanosecond,
			}).Dial,
			TLSHandshakeTimeout: time.Nanosecond,
		},
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// yields: Get "https://localhost:8888": net/http: TLS handshake timeout
