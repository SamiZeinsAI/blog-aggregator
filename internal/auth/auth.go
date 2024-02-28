package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	authSlice := strings.Split(header.Get("Authorization"), " ")
	if len(authSlice) < 2 || authSlice[0] != "ApiKey" {
		return "", errors.New("authorization header wrong shape")
	}
	return authSlice[1], nil
}
