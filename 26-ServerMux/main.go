package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/blog", blog{title: "My blog"})
	mux.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct{
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
