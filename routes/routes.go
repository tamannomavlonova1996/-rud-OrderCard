package routes

import (
	"awesomeProject2/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/cards", handlers.CreateCards).Methods(http.MethodPost)
	router.HandleFunc("/cards", handlers.GetCards).Methods("GET")
	router.HandleFunc("/cards/{id}", handlers.GetCardById).Methods("GET")
	router.HandleFunc("/cards", handlers.UpdateCardById).Methods("PUT")
	router.HandleFunc("/cards/{id}", handlers.DeleteCardById).Methods("DELETE")

	return router
}
