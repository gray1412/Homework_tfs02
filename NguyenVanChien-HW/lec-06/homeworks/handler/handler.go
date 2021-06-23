package handler

import "net/http"

type Handler interface {
	ContenTypeCheckingMiddleware(next http.Handler) http.Handler
	ContentHeaderType(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}
