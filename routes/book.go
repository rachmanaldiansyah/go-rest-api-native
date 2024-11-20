package routes

import (
	"github.com/gorilla/mux"

	"go-api-native/controllers/book_controller"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()

	router.HandleFunc("", book_controller.Index).Methods("GET")
	// router.HandleFunc("", book_controller.Create).Methods("POST")
	// router.HandleFunc("/{id}/detail", book_controller.Detail).Methods("GET")
	// router.HandleFunc("/{id}/update", book_controller.Update).Methods("PUT")
	// router.HandleFunc("/{id}/delete", book_controller.Delete).Methods("DELETE")
}
