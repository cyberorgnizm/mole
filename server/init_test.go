package server

import (
	"net/http"
	"testing"
)

func TestApplicationHandler(t *testing.T) {
	// test is expected to fail since there's no SPA ready to serve
	_, err := http.Get("localhost:8000")
	if err != nil {
		t.Log(err)
	}
}
