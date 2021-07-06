package server

import (
	"net/http"
	"tfs/tfs-api-mysql/handler"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func ServerRun() {

	routerParent := mux.NewRouter()
	router := routerParent.PathPrefix("/students").Subrouter()
	// router.Use(middleware.ContentTypeCheckingMiddleware)
	router.Methods("POST").Path("").HandlerFunc(handler.CreatePerson)
	router.Methods("GET").Path("").HandlerFunc(handler.GetAllPersons)
	router.Methods("PUT").Path("/{id:[0-9]+}").HandlerFunc(handler.UpdatePerson)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		panic(err)
	}
}
