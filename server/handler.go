package server

import (
	"encoding/json"
	"fmt"
	"github.com/zchary-ma/url-shortener/pkg"
	"log"
	"net/http"
)

var store = pkg.NewStore()

type RequestShortenInfo struct {
	URL string `json:"url"`
}

type RequestRedirectInfo struct {
	Key string `json:"key"`
}

type UniResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var info RequestShortenInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		log.Println(err)
	}

	id, err := store.Shorten(info.URL)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, fmtResponse(id))
}

func URLRedirect(w http.ResponseWriter, r *http.Request) {
	var info RequestRedirectInfo
	info.Key = r.URL.Query().Get("key")
	if info.Key == "" {
		log.Println(">>>Error: query param for key is empty.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	url, err := store.Get(info.Key)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func fmtResponse(s string) string {
	var response = &UniResponse{Code: 200, Message: "", Result: s}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	return string(bytes)
}
