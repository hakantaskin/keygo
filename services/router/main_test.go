package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	service := NewRouter()

	service.GET("/get", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {}))
	service.POST("/set", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {}))
	service.DELETE("/del", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {}))

	routeList := service.GetList()
	if routeList == nil {
		t.Errorf("route list cannot be empty")
	}

	if routeList.GET == nil {
		t.Errorf("get method list cannot be empty")
	}

	if routeList.POST == nil {
		t.Errorf("post method list cannot be empty")
	}

	if routeList.DELETE == nil {
		t.Errorf("delete method list cannot be empty")
	}

	testRequest := httptest.NewRequest("GET", "/get", nil)
	err, statusCode := service.Parser(testRequest)
	if err != nil {
		t.Error(err)
	}

	if statusCode > 500 {
		t.Errorf("status code problem")
	}

	dataByte := service.JsonResponse([]byte("ok"), nil)
	if string(dataByte) != "{\"error\":\"\",\"data\":\"ok\"}" {
		t.Errorf("got %s", string(dataByte))
	}
}