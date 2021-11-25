package service

import (
	"testing"

	"keygo/services/cache"
)

func TestKeySetterAndGetter(t *testing.T) {
	service := NewKeyService(cache.NewCacheService(cache.NewInMemoryAdapter()))
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

func TestKeyDelete(t *testing.T) {
	service := NewKeyService(cache.NewCacheService(cache.NewInMemoryAdapter()))

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

func TestKeyFlush(t *testing.T) {
	service := NewKeyService(cache.NewCacheService(cache.NewInMemoryAdapter()))

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
