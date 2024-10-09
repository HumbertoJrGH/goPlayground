package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRoot(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Couldn't create request %v", err)
	}

	rec := httptest.NewRecorder()
	GetRoot(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}

func TestGetHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	rec := httptest.NewRecorder()
	GetHello(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	expected := "Hello, Go!"
	if rec.Body.String() != expected {
		t.Errorf("Expected body to be %s, got %s", expected, rec.Body.String())
	}
}
