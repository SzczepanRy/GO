package auth

import (
	"errors"
	"net/http"
	"strings"
)

// authorization: bearer ###
func GetAPIKey(headers http.Header) (string, error) {

	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("bad auth header")
	}
	if vals[0] != "bearer" {
		return "", errors.New("bad 1 p of auth header")

	}
	return vals[1], nil
}
