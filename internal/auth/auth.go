package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetAuthorizationKey(header http.Header) (string, error) {
	authSlice := strings.Split(header.Get("Authorization"), " ")
	fmt.Printf("%s\n", authSlice[1])
	if len(authSlice) < 2 || authSlice[0] != "ApiKey" {
		return "", errors.New("authorization header wrong shape")
	}
	return authSlice[1], nil
}
