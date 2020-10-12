package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buddhamagnet/shorty/shorty"
	"github.com/buddhamagnet/shorty/store"
	"github.com/gorilla/mux"
)

const rick = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

// ID represents a short URL ID for use in responses.
type ID struct {
	LongURL string `json:"url,omitempty"`
	ShortID string `json:"id,omitempty"`
}

// APIError represents an error at the API level.
type APIError struct {
	Message string `json:"error"`
}

// ErrorResponse writes the appropriate headers and data back on error.
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
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
	var id ID
	if r.Body == nil {
		ErrorResponse(w, http.StatusBadRequest, "Empty request body")
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Error reading request")
		return
	}
	if err := json.Unmarshal(body, &id); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	if id.LongURL == "" {
		ErrorResponse(w, http.StatusBadRequest, "Please supply a URL parameter")
		return
	}

	if !shorty.IsValidURL(id.LongURL) {
		ErrorResponse(w, http.StatusBadRequest, "Please supply a valid URL")
		return
	}
	shortened, err := shorty.Shorten(id.LongURL)
	if err != nil {
		switch e := err.(type) {
		case shorty.ShortenerError:
			log.Println(e.Error())
			ErrorResponse(w, e.Code, "Unable to to store data")
			return
		default:
			log.Println(err.Error())
			ErrorResponse(w, http.StatusInternalServerError, "Unable to to store data")
			return
		}
	}
	data, _ := json.Marshal(ID{ShortID: shortened})
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
}

// Redirector redirects to a URL given a short URL ID.
func Redirector(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	longURL, err := store.Get(vars["id"])
	if err != nil || longURL == "" {
		// Rick roll.
		longURL = rick
	}
	http.Redirect(w, r, longURL, 301)
}
