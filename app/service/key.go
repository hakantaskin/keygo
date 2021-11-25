package service

import "keygo/services/cache"

type IKeyService interface {
	Set(key string, value []byte) error
	Get(key string) (error, []byte)
	Del(key string) error
	Flush() error
}

type KeyService struct {
	CacheService cache.ICache
}

func NewKeyService(cacheService cache.ICache) *KeyService {
	return &KeyService{
		CacheService: cacheService,
	}
}

func (s *KeyService) Set(key string, value []byte) error {
	return s.CacheService.Set(key, value)
}

func (s *KeyService) Get(key string) (error, []byte) {
	return s.CacheService.Get(key)
}

func (s *KeyService) Del(key string) error {
	return s.CacheService.Del(key)
}

func (s *KeyService) Flush() error {
	return s.CacheService.Flush()
}