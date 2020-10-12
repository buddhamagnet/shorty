package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ping returned incorrect status code: received %v expected %v",
			status, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("Ping returned unexpected value: received %v expected %v",
			rr.Body.String(), expected)
	}
}

func TestShortener(t *testing.T) {
	var url = []byte(`{"url":"http://www.google.com"}`)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(url))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Shortener)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Shorterne returned incorrect status code: received %v expected %v",
			status, http.StatusOK)
	}
}

func TestShortenerEmptyRequestBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/shorten", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Shortener)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Shortener returned incorrect status code: received %v expected %v",
			status, http.StatusBadRequest)
	}
}
func TestShortenerMissingURL(t *testing.T) {
	var url = []byte(`{"url":""}`)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(url))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Shortener)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Shortener returned incorrect status code: received %v expected %v",
			status, http.StatusBadRequest)
	}
}

func TestShortenerBadURL(t *testing.T) {
	var url = []byte(`{"url":"http://"}`)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(url))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Shortener)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Shortener returned incorrect status code: received %v expected %v",
			status, http.StatusBadRequest)
	}
}

func TestRedirector(t *testing.T) {
	req, err := http.NewRequest("GET", "/be4623", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"id": "be4623",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Redirector)
	handler.ServeHTTP(rr, req)

	if rr.Header().Get("Location") == "" {
		t.Errorf("Redirector returned empty Location header")
	}

	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("Redirector returned incorrect status code: received %v expected %v",
			status, http.StatusMovedPermanently)
	}
}
