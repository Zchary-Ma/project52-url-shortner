package core

import (
	"fmt"
)

type MockStorage struct {
	storage map[string]string
}

func (f MockStorage) Get(key string) (string, error) {
	if key == "ERROR" {
		return "", fmt.Errorf("error")
	}

	return f.storage[key], nil
}

func (f MockStorage) Exists(key string) (bool, error) {
	if key == "ERROR" {
		return false, fmt.Errorf("error")
	}

	_, ok := f.storage[key]
	return ok, nil
}

func (f MockStorage) Set(key string, value interface{}) error {
	if key == "ERROR" {
		return fmt.Errorf("error")
	}

	switch val := value.(type) {
	case string:
		f.storage[key] = val
		return nil
	default:
		return fmt.Errorf("Unable to store: %v", val)
	}
}

func createUrlStorage(m map[string]string) *Store {
	if m == nil {
		m = make(map[string]string)
	}

	mockStorage := MockStorage{m}
	return &Store{mockStorage}
}

// add test if necessary
