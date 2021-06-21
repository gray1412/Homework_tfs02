package handler

import "net/http"

type Handler interface {
	ContenTypeCheckingMiddleware(next http.Handler) http.Handler
	ContentHeaderType(w http.ResponseWriter, r *http.Request)
	GetStudentByName(w http.ResponseWriter, r *http.Request)
	CreateStudent(w http.ResponseWriter, r *http.Request)
	UpdateStudent(w http.ResponseWriter, r *http.Request)
	DeleteStudent(w http.ResponseWriter, r *http.Request)
}
