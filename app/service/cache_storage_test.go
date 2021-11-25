package service

import (
	"fmt"
	"testing"
	"time"

	"keygo/services/cache"
	"keygo/services/storage"
)

func TestCacheStorageGeneratePath(t *testing.T) {
	service := NewCacheStorageService(
		cache.NewCacheService(cache.NewInMemoryAdapter()),
		storage.NewStorageService(storage.NewInMemoryAdapter()))

	testFile := fmt.Sprintf("/tmp/%s-data.json", time.Now().Format("2006010215"))
	path := service.GeneratePath()

	if path != testFile {
		t.Errorf("want %s, got %s", testFile, path)
	}
}

func TestCacheStorageLoad(t *testing.T) {
	service := NewCacheStorageService(
		cache.NewCacheService(cache.NewInMemoryAdapter()),
		storage.NewStorageService(storage.NewInMemoryAdapter()))

	testFile := service.GeneratePath()
	errTestFile := fmt.Sprintf("%s-errFile", testFile)

	err := service.Load(errTestFile)
	if err == nil {
		t.Errorf("an error ocurred for %s", testFile)
	}

	err = service.Save(testFile)
	if err != nil {
		t.Error(err)
	}

	err = service.Load(testFile)
	if err != nil {
		t.Error(err)
	}
}

func TestCacheStorageSave(t *testing.T) {
	service := NewCacheStorageService(
		cache.NewCacheService(cache.NewInMemoryAdapter()),
		storage.NewStorageService(storage.NewInMemoryAdapter()))

	testFile := service.GeneratePath()
	err := service.Save(testFile)
	if err != nil {
		t.Error(err)
	}
}
