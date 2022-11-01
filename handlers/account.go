package handlers

import (
	"awesomeProject2/internal/service/account"
	"awesomeProject2/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.Account
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Code = http.StatusBadRequest
		log.Println(err)
		return
	}
	err = account.CreateAccount(req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}

	response.Message = "Данные добавились успешно!"

}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	accounts, err := account.GetAccounts()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}

	response.Message = "Данные получены успешно!"
	response.Payload = accounts
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := account.GetAccountByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные получены успешно!"
	response.Payload = result
}

func UpdateAccountByID(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.Account
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		log.Println(err)
		return
	}
	err = account.UpdateAccountByID(req)

	response.Message = "Данные обновлены успешно!"
}

func DeleteAccountByID(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := account.DeleteAccountByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные удалены успешно!"
}
