package storage

import (
	"fmt"
	"testing"
	"time"
)

func TestInMemorySetterAndGetter(t *testing.T) {
	service := NewInMemoryAdapter()
	tmpFile := fmt.Sprintf("/tmp/test_%s.json", time.Now().String())
	err := service.Set(tmpFile, []byte("155"))
	if err != nil {
		t.Error(err)
	}

	err, data := service.Get(tmpFile)
	if err != nil {
		t.Error(err)
	}

	if string(data) != "155" {
		t.Errorf("want test ok, got %s", string(data))
	}
}