package handlers

import (
	"fmt"
	"net/http"

	"buddhamagnet/shorty/shorty"
)

// Ping is the healthcheck endpoint for the service.
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// Shortener is the URL shortener endpoint.
func Shortener(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please supply a URL parameter")
		return
	}

	if !shorty.IsValidURL(longURL) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please supply a valid URL")
		return
	}
	fmt.Fprintf(w, shorty.Shorten())
}

// Decoder redirects to a URL given a short
func Decoder(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.google.com", 301)
}
