package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := strings.TrimSpace(headers.Get("Authorization"))
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", fmt.Errorf("invalid authorization header")
	}

	tok := strings.TrimSpace(strings.TrimPrefix(authHeader, bearerPrefix))
	if tok == "" {
		return "", fmt.Errorf("empty bearer token")
	}

	return tok, nil
}
