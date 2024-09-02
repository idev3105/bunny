package tokenutil

import (
	"net/http"
	"strings"
)

// GetTokenFromHeader extracts the token from the Authorization header
func GetTokenFromHeader(request *http.Request) string {
	authHeader := request.Header.Get("Authorization")
	return strings.TrimPrefix(authHeader, "Bearer ")
}

// GetTokenFromCookies retrieves the token from the "token" cookie
func GetTokenFromCookies(request *http.Request) string {
	cookie, _ := request.Cookie("token")
	if cookie != nil {
		return cookie.Value
	}
	return ""
}
