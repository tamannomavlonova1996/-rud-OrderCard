package handlers

import (
	user2 "awesomeProject2/internal/user"
	"awesomeProject2/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		user     user2.User
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "Неверные данные"
		log.Println(err)
		return
	}
	err = user.CreateUser()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}

	response.Message = "Данные добавились успешно!"
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var (
		user     user2.User
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Code = http.StatusBadRequest
		log.Println(err)
		return
	}
	users, err := user.GetUsers()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}

	response.Message = "Данные получены успешно!"
	response.Payload = users
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var (
		user     user2.User
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := user.GetUserByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные получены успешно!"
	response.Payload = result
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	var (
		user     user2.User
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Code = http.StatusInternalServerError
		log.Println(err)
		return
	}
	err = user.UpdateUserByID()

	response.Message = "Данные обновлены успешно!"
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	var (
		user     user2.User
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := user.DeleteUserByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные удалены успешно!"
}
