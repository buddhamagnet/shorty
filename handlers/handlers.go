package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/buddhamagnet/shorty/shorty"
	"github.com/buddhamagnet/shorty/store"
	"github.com/gorilla/mux"
)

const rick = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

// ID represents a short URL ID for use in responses.
type ID struct {
	ShortID string `json:"id"`
}

// APIError represents an error at the API level.
type APIError struct {
	Message string `json:"error"`
}

func ErrorReponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	err, _ := json.Marshal(APIError{message})
	fmt.Fprintf(w, string(err))
}

// Ping is the healthcheck endpoint for the service.
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// Shortener is the URL shortener endpoint.
func Shortener(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		ErrorReponse(w, http.StatusBadRequest, "Please supply a URL parameter")
		return
	}

	if !shorty.IsValidURL(longURL) {
		ErrorReponse(w, http.StatusBadRequest, "Please supply a valid URL")
		return
	}
	shortened, err := shorty.Shorten(longURL)
	if err != nil {
		switch e := err.(type) {
		case shorty.ShortenerError:
			log.Println(e.Error())
			ErrorReponse(w, e.Code, "Unable to to store data")
			return
		default:
			log.Println(err.Error())
			ErrorReponse(w, http.StatusInternalServerError, "Unable to to store data")
			return
		}
	}
	data, _ := json.Marshal(ID{shortened})
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
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
