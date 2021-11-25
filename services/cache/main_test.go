package cache

import (
	"testing"
)

func TestSetterAndGetter(t *testing.T) {
	service := NewInMemoryAdapter()
	err := service.Set("test_1", []byte("155"))
	if err != nil {
		t.Error(err)
	}

	err, data := service.Get("test_1")
	if err != nil {
		t.Error(err)
	}

	if string(data) != "155" {
		t.Errorf("want test ok, got %s", string(data))
	}
}
