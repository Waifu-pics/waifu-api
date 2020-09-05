package web

import (
	"errors"
	"net/http"
)

func GetCookie(cookies []*http.Cookie, key string) (string, error) {
	for _, cookie := range cookies {
		if cookie.Name == key {
			return cookie.Value, nil
		}
	}
	return "", errors.New("there was no such cookie: " + key)
}
