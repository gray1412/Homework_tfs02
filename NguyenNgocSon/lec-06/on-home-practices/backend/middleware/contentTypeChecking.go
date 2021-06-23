package middleware

import (
	"fmt"
	"net/http"
)

func ContentTypeChecking(next http.Handler) http.Handler {
	const contentType = "application/json;charset=UTF-8"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != contentType {
			fmt.Fprintf(w, "Only allow request with content type %v", contentType)
			return
		}
		next.ServeHTTP(w, r)
	})
}
