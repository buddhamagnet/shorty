package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
	req, err := http.NewRequest("GET", "/shorten?url=http://google.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Shortener)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Decoder returned incorrect status code: received %v expected %v",
			status, http.StatusOK)
	}
}

func TestShortenerMissingURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/shorten", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Shortener)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Decoder returned incorrect status code: received %v expected %v",
			status, http.StatusBadRequest)
	}

	expected := "Please supply a URL parameter"
	if rr.Body.String() != expected {
		t.Errorf("Ping returned unexpected value: received %v expected %v",
			rr.Body.String(), expected)
	}
}

func TestShortenerBadURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/shorten?url=http://", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Shortener)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Decoder returned incorrect status code: received %v expected %v",
			status, http.StatusBadRequest)
	}

	expected := "Please supply a valid URL"
	if rr.Body.String() != expected {
		t.Errorf("Ping returned unexpected value: received %v expected %v",
			rr.Body.String(), expected)
	}
}

func TestDecoder(t *testing.T) {
	req, err := http.NewRequest("GET", "/decimate", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Decoder)
	handler.ServeHTTP(rr, req)

	if rr.Header().Get("Location") == "" {
		t.Errorf("Decoder returned empty Location header")
	}

	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("Decoder returned incorrect status code: received %v expected %v",
			status, http.StatusMovedPermanently)
	}
}

type URLTest struct {
	url      string
	expected bool
}
