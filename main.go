package main

import (
	"github.com/zchary-ma/url-shortener/server"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/shorten", server.ShortenURL)
	_ = http.ListenAndServe(":8080", r)
}
