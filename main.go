package main

import (
	"github.com/zchary-ma/url-shortener/server"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/shorten", server.ShortenURL)
	r.HandleFunc("/redirect", server.URLRedirect)
	_ = http.ListenAndServe(":8080", r)
}
