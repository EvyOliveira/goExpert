package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCEPHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipCodeParam := r.URL.Query().Get("cep")
	if zipCodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
