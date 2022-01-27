package storage

import "errors"

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type KvStorage struct {
	Store map[string]string
}

func NewKvStorage() *KvStorage {
	s := make(map[string]string, 100)
	return &KvStorage{Store: s}
}

func (s *KvStorage) Init() {
	s.Store = make(map[string]string, 100)
}

func (s *KvStorage) Get(key string) (string, error) {
	if val, ok := s.Store[key]; ok {
		return val, nil
	}
	return "", errors.New("key not found")
}

func (s *KvStorage) Set(key string, value string) error {
	s.Store[key] = value
	return nil
}
