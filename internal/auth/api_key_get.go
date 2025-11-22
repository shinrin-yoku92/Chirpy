package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := strings.TrimSpace(headers.Get("Authorization"))
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	const keyPrefix = "ApiKey "
	if !strings.HasPrefix(authHeader, keyPrefix) {
		return "", fmt.Errorf("invalid authorization header")
	}

	key := strings.TrimSpace(strings.TrimPrefix(authHeader, keyPrefix))
	if key == "" {
		return "", fmt.Errorf("empty API key")
	}

	return key, nil
}
