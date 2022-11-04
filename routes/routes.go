package routes

import (
	"awesomeProject2/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/auth/sign-up", handlers.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/auth/sign-in", handlers.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/auth/changePassword", handlers.ChangePassword).Methods(http.MethodPut)
	router.HandleFunc("/auth/resetPassword", handlers.ResetPassword).Methods(http.MethodPut)

	routerApi := router.PathPrefix("").Subrouter()
	routerApi.Use(handlers.AuthorizeMiddleware)

	//order card
	routerApi.HandleFunc("/orderCard", handlers.CreateOrderCard).Methods(http.MethodPost)
	routerApi.HandleFunc("/orderCard", handlers.GetOrderCards).Methods(http.MethodGet)
	routerApi.HandleFunc("/orderCard/{id}", handlers.GetOrderCardByID).Methods(http.MethodGet)
	routerApi.HandleFunc("/orderCard", handlers.UpdateOrderCardByID).Methods(http.MethodPut)
	routerApi.HandleFunc("/orderCard/{id}", handlers.DeleteOrderCardByID).Methods(http.MethodDelete)

	//user
	routerApi.HandleFunc("/user", handlers.CreateUser).Methods(http.MethodPost)
	routerApi.HandleFunc("/user", handlers.GetUsers).Methods(http.MethodGet)
	routerApi.HandleFunc("/user/{id}", handlers.GetUserByID).Methods(http.MethodGet)
	routerApi.HandleFunc("/user", handlers.UpdateUserByID).Methods(http.MethodPut)
	routerApi.HandleFunc("/user/{id}", handlers.DeleteUserByID).Methods(http.MethodDelete)

	//card
	routerApi.HandleFunc("/card", handlers.CreateCard).Methods(http.MethodPost)
	routerApi.HandleFunc("/card", handlers.GetCards).Methods(http.MethodGet)
	routerApi.HandleFunc("/card/{id}", handlers.GetCardByID).Methods(http.MethodGet)
	routerApi.HandleFunc("/card", handlers.UpdateCardByID).Methods(http.MethodPut)
	routerApi.HandleFunc("/card/{id}", handlers.DeleteCardByID).Methods(http.MethodDelete)

	//account
	routerApi.HandleFunc("/account", handlers.CreateAccount).Methods(http.MethodPost)
	routerApi.HandleFunc("/account", handlers.GetAccounts).Methods(http.MethodGet)
	routerApi.HandleFunc("/account/{id}", handlers.GetAccountByID).Methods(http.MethodGet)
	routerApi.HandleFunc("/account", handlers.UpdateAccountByID).Methods(http.MethodPut)
	routerApi.HandleFunc("/account/{id}", handlers.DeleteAccountByID).Methods(http.MethodDelete)

	//operation
	routerApi.HandleFunc("/operation", handlers.CreateOperation).Methods(http.MethodPost)
	routerApi.HandleFunc("/operation", handlers.GetOperations).Methods(http.MethodGet)
	routerApi.HandleFunc("/operation/{id}", handlers.GetOperationByID).Methods(http.MethodGet)
	routerApi.HandleFunc("/operation", handlers.UpdateOperationByID).Methods(http.MethodPut)
	routerApi.HandleFunc("/operation/{id}", handlers.DeleteOperationByID).Methods(http.MethodDelete)

	return router
}
