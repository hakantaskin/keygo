package cache

type ICache interface {
	Set(key string, value []byte) error
	Get(key string) (error, []byte)
	Del(key string) error
	Flush() error
	GetAllData() []byte
	LoadData([]byte) error
}

type Service struct {
	Adapter ICache
}

func NewCacheService(service ICache) *Service {
	return &Service{Adapter: service}
}

func (s *Service) Set(key string, value []byte) error {
	return s.Adapter.Set(key, value)
}

func (s *Service) Get(key string) (error, []byte) {
	return s.Adapter.Get(key)
}

func (s *Service) Del(key string) error {
	return s.Adapter.Del(key)
}

func (s *Service) Flush() error {
	return s.Adapter.Flush()
}

func (s *Service) GetAllData() []byte  {
	return s.Adapter.GetAllData()
}

func (s *Service) LoadData(data []byte) error {
	return s.Adapter.LoadData(data)
}