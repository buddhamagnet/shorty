package shorty

import (
	"net/url"

	"github.com/buddhamagnet/shorty/store"
	"github.com/google/uuid"
)

// Shorten generates a 6-character string from a UUID.
func Shorten(url string) (string, error) {
	shortened := uuid.New().String()[:6]
	return store.Put(url, shortened)
}

// Decode returns the long URL for a short URL.
func Decode(shortened string) (string, error) {
	return store.Get(shortened)
}

// IsValidURL determines whether a URL is valid for shortening. The
// stdlib functions (url.ParseRequestURI etc) are lacking in this
// respect so we roll our own.
func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
