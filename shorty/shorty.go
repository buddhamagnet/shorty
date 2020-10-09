package shorty

import (
	"net/url"

	"github.com/google/uuid"
)

// Shorten generates a 6-character string from a UUID.
func Shorten() string {
	return uuid.New().String()[:6]
}

// IsValidURL determines whether a URL is valid for shortening. The
// stdlib functions (url.ParseRequestURI etc) are lacking in this
// respect so we roll our own.
func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
