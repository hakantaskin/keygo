package cache

import (
	"encoding/json"
	"fmt"
	"sync"
)

type InMemory struct {
	mu sync.Mutex
	cacheList map[string]interface{}
}

func NewInMemoryAdapter() ICache {
	return &InMemory{cacheList: make(map[string]interface{})}
}

func (s *InMemory) Set(key string, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var data interface{}
	err := json.Unmarshal(value, &data)
	if err != nil {
		data = string(value)
	}

	s.cacheList[key] = data

	return nil
}

func (s *InMemory) Get(key string) (error, []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cacheList[key] == nil {
		return fmt.Errorf("record not found! key: %s", key), nil
	}

	j, err := json.Marshal(s.cacheList[key])
	if err != nil {
		return err, nil
	}

	return nil, j
}

func (s *InMemory) Del(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cacheList[key] == nil {
		return fmt.Errorf("record not found! key: %s", key)
	}

	delete(s.cacheList, key)
	return nil
}

func (s *InMemory) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cacheList = make(map[string]interface{})
	return nil
}

func (s *InMemory) GetAllData() []byte {
	j, err := json.Marshal(s.cacheList)
	if err != nil {
		return nil
	}
	return j
}

func (s *InMemory) LoadData(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var loadData map[string]interface{}
	err := json.Unmarshal(data, &loadData)
	if err != nil {
		return err
	}

	s.cacheList = loadData
	return nil
}