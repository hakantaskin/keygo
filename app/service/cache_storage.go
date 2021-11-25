package service

import (
	"fmt"
	"time"

	"keygo/services/cache"
	"keygo/services/storage"
)

type ICacheStorage interface {
	Load(path string) error
	Save(path string) error
	GeneratePath() string
	Run()
}

type CacheStorageService struct {
	CacheService cache.ICache
	StorageService storage.IStorage
}

func NewCacheStorageService(cacheService cache.ICache, storage storage.IStorage) *CacheStorageService {
	return &CacheStorageService{
		CacheService: cacheService,
		StorageService: storage,
	}
}

func (s *CacheStorageService) GeneratePath() string {
	return fmt.Sprintf("/tmp/%s-data.json", time.Now().Format("2006010215"))
}

func (s *CacheStorageService) Run() {
	err := s.Load(s.GeneratePath())
	if err != nil {
		fmt.Println(err)
	}

	for {
		time.Sleep(10 * time.Second)
		err = s.Save(s.GeneratePath())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *CacheStorageService) Load(path string) error {
	err, data := s.StorageService.Get(path)
	if err != nil {
		return err
	}

	return s.CacheService.LoadData(data)
}

func (s *CacheStorageService) Save(path string) error {
	return s.StorageService.Set(path, s.CacheService.GetAllData())
}