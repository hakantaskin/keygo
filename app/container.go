package app

import (
	"keygo/app/controller"
	"keygo/app/service"
	"keygo/services/cache"
	"keygo/services/router"
	"keygo/services/storage"
)

// ContainerServicesCache /**
var (
	ContainerServicesCache = cache.NewCacheService(cache.NewInMemoryAdapter())
	ContainerRouterServices = router.NewRouter()
	ContainerServicesStorage = storage.NewStorageService(storage.NewInMemoryAdapter())
)

// Internal Services /**
var (
	ContainerServiceKey = service.NewKeyService(ContainerServicesCache)
	ContainerServiceCacheStorage = service.NewCacheStorageService(ContainerServicesCache, ContainerServicesStorage)
)

// ContainerKeyController /**
var (
	ContainerKeyController = controller.NewKeyController(ContainerRouterServices, ContainerServiceKey)
)