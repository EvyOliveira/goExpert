package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com/")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}
