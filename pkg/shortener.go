package pkg

import (
	"fmt"
	"github.com/teris-io/shortid"
)

type Shortener interface {
	Shorten(url string) (string, error)
	Get(key string) (string, error)
}

type Store struct {
	storage Storage
}

func New(s Storage) *Store {
	return &Store{s}
}

func generateId() (string, error) {
	return shortid.Generate()
}

func (s *Store) shorten(url string) (string, error) {
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
