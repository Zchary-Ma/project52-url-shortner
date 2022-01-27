package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zchary-ma/url-shortener/storage"
	"net/http"
)

var resource = storage.NewResource()

const hashLength = 5

type ShortenReq struct {
	OriginUrl string `json:"url"`
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	shortened := mux.Vars(r)["shortened"]
	if shortened == "" {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	originUrl, err := resource.Storage.Get(shortened)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, originUrl, http.StatusFound)
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	req := &ShortenReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shortened, err := resource.Storage.Get(req.OriginUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if shortened == "" {
		shortened = hash(req.OriginUrl, hashLength)
		err := resource.Storage.Set(shortened, req.OriginUrl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte(shortened))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("unknown"))
}

func hash(str string, length int) (hashed string) {
	hashed = str[:length]
	return hashed
}
