package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"keygo/app/service"
	"keygo/services/router"
)

type IKeyController interface {
	Set(key string, value interface{}) error
	Get(key string) (error, interface{})
	Del(key string) error
	Flush() error
}

type KeyController struct {
	RouteService router.IRouter
	KeyService service.IKeyService
}

func NewKeyController(r router.IRouter, keyService service.IKeyService) *KeyController {
	return &KeyController{RouteService: r, KeyService: keyService}
}


func (s *KeyController) Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err, statusCode := s.RouteService.Parser(r)
	if err != nil {
		w.WriteHeader(statusCode)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	keyParam := r.URL.Query().Get("key")
	if keyParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, fmt.Errorf("missing key")))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	err = s.KeyService.Set(keyParam, body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(s.RouteService.JsonResponse(body, nil))
	return
}

func (s *KeyController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err, statusCode := s.RouteService.Parser(r)
	if err != nil {
		w.WriteHeader(statusCode)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	keyParam := r.URL.Query().Get("key")
	if keyParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, fmt.Errorf("missing key")))
		return
	}

	err, data := s.KeyService.Get(keyParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(s.RouteService.JsonResponse(data, nil))
}

func (s *KeyController) Del(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err, statusCode := s.RouteService.Parser(r)
	if err != nil {
		w.WriteHeader(statusCode)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	keyParam := r.URL.Query().Get("key")
	if keyParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, fmt.Errorf("missing key")))
		return
	}

	err = s.KeyService.Del(keyParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(s.RouteService.JsonResponse([]byte("ok"), nil))
}

func (s *KeyController) Flush(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err, statusCode := s.RouteService.Parser(r)
	if err != nil {
		w.WriteHeader(statusCode)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	err = s.KeyService.Flush()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(s.RouteService.JsonResponse(nil, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(s.RouteService.JsonResponse([]byte("ok"), nil))
}