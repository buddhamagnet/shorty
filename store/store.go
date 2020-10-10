package store

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

func dial() (redis.Conn, error) {
	return redis.Dial("tcp", "localhost:6379")
}

// Put persists data in the back end.
func Put(url, shortened string) (string, error) {
	conn, err := dial()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	_, err = conn.Do("SET", shortened, url)
	if err != nil {
		return "", err
	}
	return shortened, err
}

// Get retrieves the value for a short key.
func Get(shortened string) (string, error) {
	conn, err := dial()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	longURL, err := redis.String(conn.Do("GET", shortened))
	if err != nil {
		return "", errors.New("unable to retrieve content")
	}
	return longURL, err
}
