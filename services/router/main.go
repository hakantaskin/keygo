package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Error string `json:"error"`
	Data interface{} `json:"data"`
}

type (
	RouteList struct {
		GET map[string]http.HandlerFunc
		POST map[string]http.HandlerFunc
		DELETE map[string]http.HandlerFunc
	}

	Router struct {
		RouteList *RouteList
	}

	IRouter interface {
		GET(prefix string, handler http.HandlerFunc)
		POST(prefix string, handler http.HandlerFunc)
		DELETE(prefix string, handler http.HandlerFunc)
		GetList() *RouteList
		Parser(request *http.Request) (error, int)
		JsonResponse(value []byte, err error) []byte
	}
)

func NewRouter() IRouter {
	return &Router{RouteList: &RouteList{
		GET: make(map[string]http.HandlerFunc),
		POST: make(map[string]http.HandlerFunc),
		DELETE: make(map[string]http.HandlerFunc),
	}}
}

func (r *Router) GET(prefix string, handler http.HandlerFunc) {
	r.add(http.MethodGet, prefix, handler)
}

func (r *Router) POST(prefix string, handler http.HandlerFunc) {
	r.add(http.MethodPost, prefix, handler)
}

func (r *Router) DELETE(prefix string, handler http.HandlerFunc) {
	r.add(http.MethodDelete, prefix, handler)
}

func (r *Router) GetList() *RouteList  {
	return r.RouteList
}

func (r *Router) add(method string, prefix string, handler http.HandlerFunc)  {
	switch method {
	case http.MethodGet:
		r.RouteList.GET[prefix] = handler
	case http.MethodPost:
		r.RouteList.POST[prefix] = handler
	case http.MethodDelete:
		r.RouteList.DELETE[prefix] = handler
	}
}

func (r *Router) Parser(request *http.Request) (error, int) {
	prefix := request.URL.Path

	switch request.Method {
	case http.MethodGet:
		if r.RouteList.GET[prefix] == nil {
			return fmt.Errorf("prefix not found in get list %s", request.RequestURI), http.StatusNotFound
		}
	case http.MethodPost:
		if r.RouteList.POST[prefix] == nil {
			return fmt.Errorf("prefix not found in post list %s", request.RequestURI), http.StatusNotFound
		}
	case http.MethodDelete:
		if r.RouteList.DELETE[prefix] == nil {
			return fmt.Errorf("prefix not found in delete list %s", request.RequestURI), http.StatusNotFound
		}
	default:
		return fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed
	}

	return nil, http.StatusOK
}

func (r *Router) JsonResponse(value []byte, err error) []byte {
	var errorString string
	if err != nil {
		errorString = err.Error()
	}

	var jsonData interface{}
	if value != nil {
		err = json.Unmarshal(value, &jsonData)
		if err != nil {
			jsonData = string(value)
		}
	}

	js, err := json.Marshal(JsonResponse{Data: jsonData, Error: errorString})
	if err != nil {
		return nil
	}

	return js
}