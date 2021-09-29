package pkg

import (
	"fmt"
	"github.com/teris-io/shortid"
)

type IShortener interface {
	Shorten(url string) (string, error)
	Get(key string) (string, error)
}

// Store NOTE: how to initiate struct with preconfigure
type Store struct {
	storage Storage
}

func NewStore() Store {
	s := Store{}
	var configurations Configurations
	LoadConfig(&configurations)
	s.storage = CreateClient(configurations)
	return s
}

func generateId() (string, error) {
	return shortid.Generate()
}

func (s *Store) Shorten(url string) (string, error) {
	id, err := generateId()
	if err != nil {
		return "", fmt.Errorf("generating ID error: %s", err)
	}

	ok, err := s.storage.Exists(id)
	if err != nil {
		return "", fmt.Errorf("checking existance of key %s for url %s: %v", id, url, err)
	} else if ok {
		return "", fmt.Errorf("key %s for url %s already exists in storage", id, url)
	}

	if err := s.storage.Set(id, url); err != nil {
		return "", fmt.Errorf("storing key %s for url %s: %v", id, url, err)
	}

	return id, nil
}

func (s *Store) Get(key string) (string, error) {
	return s.storage.Get(key)
}
