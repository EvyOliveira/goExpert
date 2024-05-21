package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request started")
	defer log.Println("request completed")

	select {
	case <-time.After(5 * time.Second):
		//print to stdout command line
		log.Println("request processed successfully")
		//print in browser
		w.Write([]byte("request processed successfully"))
	case <-ctx.Done():
		//print to stdout command line
		log.Println("request canceled by client")
	}
}
