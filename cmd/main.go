package main

import (
	"github.com/gorilla/mux"
	"github.com/zchary-ma/url-shortener/server"
	"log"
	"net/http"
)

var port = "8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("url shortening server"))
	})
	r.HandleFunc("/api/h", server.HealthCheck)
	r.HandleFunc("/api/s", server.Shorten)
	r.HandleFunc("/api/{shortened}", server.Redirect)

	go func() {
		log.Printf("server is running on port: %s", port)
	}()

	log.Fatal(http.ListenAndServe(":"+port, r))
}
