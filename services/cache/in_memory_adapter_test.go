package cache

import (
	"encoding/json"
	"testing"
)

func TestInMemorySetterAndGetter(t *testing.T) {
	service := NewInMemoryAdapter()
	err := service.Set("test_1", []byte("155"))
	if err != nil {
		t.Error(err)
	}

	err, data := service.Get("test_1")
	if err != nil {
		t.Error(err)
	}

	if string(data) != "155" {
		t.Errorf("want test ok, got %s", string(data))
	}
}

func TestInMemoryDelete(t *testing.T) {
	service := NewInMemoryAdapter()

	err := service.Set("test_1", []byte("155"))
	if err != nil {
		t.Error(err)
	}

	err = service.Del("test_1")
	if err != nil {
		t.Error(err)
	}

	err, _ = service.Get("test_1")
	if err == nil {
		t.Errorf("key delete problem")
	}
}

func TestInMemoryFlush(t *testing.T) {
	service := NewInMemoryAdapter()

	err := service.Set("test_1", []byte("155"))
	if err != nil {
		t.Error(err)
	}

	err = service.Flush()
	if err != nil {
		t.Error(err)
	}

	err, _ = service.Get("test_1")
	if err == nil {
		t.Errorf("key delete problem")
	}
}

func TestInMemoryGetAllData(t *testing.T) {
	service := NewInMemoryAdapter()
	if service.GetAllData() == nil {
		t.Errorf("cache list cannot be empty")
	}
}

func (s *InMemory) TestInMemoryLoadData(t *testing.T) {
	service := NewInMemoryAdapter()
	data := map[string]interface{}{
		"deneme": 911,
	}

	j, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	err = service.LoadData(j)
	if err != nil {
		t.Error(err)
	}
}