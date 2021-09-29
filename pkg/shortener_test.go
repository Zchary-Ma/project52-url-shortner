package pkg

import (
	"errors"
	"fmt"
	"testing"
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
func TestShortenGenerateID(t *testing.T) {
	tests := []struct {
		key         string
		generateID  func() (string, error)
		shouldError bool
	}{
		{"should succeed", func() (string, error) { return "KEY", nil }, false},
		{"should error", func() (string, error) { return "", errors.New("Error") }, true},
	}
	idGenerator := generateId
	for _, test := range tests {
		store := createURLStorage(nil)
		generateID = test.generateID

		_, err := store.Shorten(test.key)
		if test.shouldError && err == nil {
			t.Errorf("store.Shorten(%s) expected error, found none", test.key)
		}

		if !test.shouldError && err != nil {
			t.Errorf("store.Shorten(%s) found error, expected none: %v", test.key, err)
		}

	}

}
