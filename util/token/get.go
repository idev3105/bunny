package tokenutil

import "net/http"

func GetTokenFromHeader(request http.Request) string {
	return request.Header.Get("Authorization")[len("Bearer "):]
}

func GetTokenFromCookies(request http.Request) string {
	cookie, err := request.Cookie("token")
	if err != nil {
		return ""
	}
	return cookie.Value
}
