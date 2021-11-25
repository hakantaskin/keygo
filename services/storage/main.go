package storage

type IStorage interface {
	Set(path string, data []byte) error
	Get(path string) (error, []byte)
}

type Service struct {
	Adapter IStorage
}

func NewStorageService(service IStorage) *Service {
	return &Service{Adapter: service}
}

func (s *Service) Set(path string, data []byte) error {
	return s.Adapter.Set(path, data)
}

func (s *Service) Get(key string) (error, []byte) {
	return s.Adapter.Get(key)
}
