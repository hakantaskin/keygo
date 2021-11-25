package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"keygo/app/service"
	"keygo/services/cache"
	"keygo/services/router"
)

func TestKeyController(t *testing.T) {
	routerService := router.NewRouter()
	cacheService := cache.NewCacheService(cache.NewInMemoryAdapter())
	keyService := service.NewKeyService(cacheService)

	keyController := NewKeyController(routerService, keyService)

	t.Run("set", func(t *testing.T) {
		routerService.POST("/set", keyController.Set)

		req := httptest.NewRequest(http.MethodPost, "/set", bytes.NewBuffer([]byte("911")))
		q := url.Values{}
		q.Add("key", "deneme")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		keyController.Set(w, req)
		resp := w.Result()

		body, _ := io.ReadAll(resp.Body)
		var response router.JsonResponse
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("Error unmarshalling response: %s", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code error %d", resp.StatusCode)
		}

		if  response.Data != float64(911) {
			t.Errorf("want %d, got %d", response.Data, 911)
		}
	})

	t.Run("get", func(t *testing.T) {
		routerService.GET("/get", keyController.Get)

		req := httptest.NewRequest(http.MethodGet, "/get", nil)
		q := url.Values{}
		q.Add("key", "deneme")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		keyController.Get(w, req)
		resp := w.Result()

		body, _ := io.ReadAll(resp.Body)
		var response router.JsonResponse
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("Error unmarshalling response: %s", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code error %d", resp.StatusCode)
		}

		if  response.Data != float64(911) {
			t.Errorf("want %d, got %d", response.Data, 911)
		}
	})

	t.Run("del", func(t *testing.T) {
		routerService.DELETE("/del", keyController.Del)

		req := httptest.NewRequest(http.MethodDelete, "/del", nil)
		q := url.Values{}
		q.Add("key", "deneme")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		keyController.Del(w, req)
		resp := w.Result()

		body, _ := io.ReadAll(resp.Body)
		var response router.JsonResponse
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("Error unmarshalling response: %s", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code error %d", resp.StatusCode)
		}
	})

	t.Run("flush", func(t *testing.T) {
		routerService.DELETE("/flush", keyController.Flush)

		req := httptest.NewRequest(http.MethodDelete, "/flush", nil)
		q := url.Values{}
		q.Add("key", "deneme")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		keyController.Flush(w, req)
		resp := w.Result()

		body, _ := io.ReadAll(resp.Body)
		var response router.JsonResponse
		err := json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("Error unmarshalling response: %s", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code error %d", resp.StatusCode)
		}
	})
}