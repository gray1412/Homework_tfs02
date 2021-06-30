package router

import (
	"amazon/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Search() {
	r := mux.NewRouter().StrictSlash(true)
	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/search").HandlerFunc(search)
	get.Use(contentTypeJson)
	http.Handle("/", r)
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
	}).Handler(r)
	// log.Fatal(http.ListenAndServe(":8082", handler))
	http.ListenAndServe(":8000", handler)
}
func search(w http.ResponseWriter, r *http.Request) {
	db := storage.Connect()
	var s storage.Review
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var review []storage.Review
	start := time.Now()
	db.Where("body LIKE ?", "%"+s.Body+"%").Find(&review)
	end := time.Since(start)
	fmt.Print(end)
	b, _ := json.Marshal(review)
	fmt.Fprint(w, string(b))
	defer db.Close()
}
func contentTypeJson(next http.Handler) http.Handler {
	const contentType = "application/json"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		if reqContentType != contentType {
			fmt.Fprintf(w, "Only allow request with content type %v", contentType)
			return
		}
		next.ServeHTTP(w, r)
	})
}
