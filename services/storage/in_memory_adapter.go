package storage

import (
	"os"
	"sync"
)

type InMemory struct {
	mu sync.Mutex
}

func NewInMemoryAdapter() IStorage {
	return &InMemory{}
}

func (s *InMemory) Set(path string, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *InMemory) Get(path string) (error, []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(path)
	if err != nil {
		return err, nil
	}

	return nil, data
}
