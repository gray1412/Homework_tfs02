package middleware

import (
	"net/http"
	"tfs/tfs-api/handlers"
)

const jsonContentType = "application/json"

func ContentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != jsonContentType {
			handlers.ResponseWithJson(w, map[string]string{"message": "request only allow content type application/json"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
