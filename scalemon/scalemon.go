package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler)

	err := http.ListenAndServe(":8085", nil)

	if err != nil {
		fmt.Printf("error happened: %s\n", err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
}
