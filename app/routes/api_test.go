package routes

import (
	"testing"

	"keygo/services/router"
)

func TestDefineApiRoute(t *testing.T) {
	routerService := router.NewRouter()
	DefineApiRoute(routerService)

	routeList := routerService.GetList()
	if routeList.GET == nil {
		t.Errorf("router get list cannot be empty")
	}

	if routeList.POST == nil {
		t.Errorf("router post list cannot be empty")
	}

	if routeList.DELETE == nil {
		t.Errorf("router delete list cannot be empty")
	}
}
