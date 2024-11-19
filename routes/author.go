package routes

import (
	"github.com/gorilla/mux"

	"go-api-native/controllers"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", controllers.Index).Methods("GET")
	router.HandleFunc("", controllers.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", controllers.Detail).Methods("GET")
}
