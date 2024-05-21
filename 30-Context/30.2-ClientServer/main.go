package main

import "net/http"

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
