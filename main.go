package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port = "8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("url shortening server"))
	})

	go func() {
		log.Printf("server is running on port: %s", port)
	}()

	http.ListenAndServe(":"+port, r)
}
