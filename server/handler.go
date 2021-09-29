package server

import (
	"encoding/json"
	"fmt"
	"github.com/zchary-ma/url-shortener/pkg"
	"log"
	"net/http"
)

type RequestInfo struct {
	URL string `json:"url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var (
		info  RequestInfo
		store pkg.Store = pkg.NewStore()
	)
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		log.Println(err)
	}

	uuid, err := store.Shorten(info.URL)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "uuid is %s", uuid)
}

func URLRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
