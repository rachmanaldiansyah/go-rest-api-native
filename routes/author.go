package routes

import (
	"github.com/gorilla/mux"

	"go-api-native/controllers/author_controller"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", author_controller.Index).Methods("GET")
	router.HandleFunc("", author_controller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", author_controller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", author_controller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", author_controller.Delete).Methods("DELETE")
}
