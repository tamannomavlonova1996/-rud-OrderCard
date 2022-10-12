package routes

import (
	"awesomeProject2/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routers() *mux.Router {
	router := mux.NewRouter()

	//order card
	router.HandleFunc("/orderCard", handlers.CreateOrderCard).Methods(http.MethodPost)
	router.HandleFunc("/orderCard", handlers.GetOrderCards).Methods(http.MethodGet)
	router.HandleFunc("/orderCard/{id}", handlers.GetOrderCardByID).Methods(http.MethodGet)
	router.HandleFunc("/orderCard", handlers.UpdateOrderCardByID).Methods(http.MethodPut)
	router.HandleFunc("/orderCard/{id}", handlers.DeleteOrderCardByID).Methods(http.MethodDelete)

	//user
	router.HandleFunc("/user", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", handlers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", handlers.GetUserByID).Methods(http.MethodGet)
	router.HandleFunc("/user", handlers.UpdateUserByID).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", handlers.DeleteUserByID).Methods(http.MethodDelete)

	//card
	router.HandleFunc("/orderCard", handlers.CreateCard).Methods(http.MethodPost)
	router.HandleFunc("/orderCard", handlers.GetCards).Methods(http.MethodGet)
	router.HandleFunc("/orderCard/{id}", handlers.GetCardByID).Methods(http.MethodGet)
	router.HandleFunc("/orderCard", handlers.UpdateCardByID).Methods(http.MethodPut)
	router.HandleFunc("/orderCard/{id}", handlers.DeleteCardByID).Methods(http.MethodDelete)

	//account
	router.HandleFunc("/account", handlers.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/account", handlers.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/account/{id}", handlers.GetAccountByID).Methods(http.MethodGet)
	router.HandleFunc("/account", handlers.UpdateAccountByID).Methods(http.MethodPut)
	router.HandleFunc("/account/{id}", handlers.DeleteAccountByID).Methods(http.MethodDelete)

	//operation
	router.HandleFunc("/operation", handlers.CreateOperation).Methods(http.MethodPost)
	router.HandleFunc("/operation", handlers.GetOperations).Methods(http.MethodGet)
	router.HandleFunc("/operation/{id}", handlers.GetOperationByID).Methods(http.MethodGet)
	router.HandleFunc("/operation", handlers.UpdateOperationByID).Methods(http.MethodPut)
	router.HandleFunc("/operation/{id}", handlers.DeleteOperationByID).Methods(http.MethodDelete)

	return router
}
