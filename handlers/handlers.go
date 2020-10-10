package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/buddhamagnet/shorty/shorty"
	"github.com/buddhamagnet/shorty/store"
	"github.com/gorilla/mux"
)

const rick = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

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
	shortened, err := shorty.Shorten(longURL)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to store data")
		return

	}
	fmt.Fprintf(w, shortened)
}

// Redirector redirects to a URL given a short
func Redirector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	log.Println(vars)
	longURL, err := store.Get(vars["id"])
	if err != nil || longURL == "" {
		// Rick roll.
		longURL = rick
	}
	http.Redirect(w, r, longURL, 301)
}
