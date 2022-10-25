package handlers

import (
	"awesomeProject2/helpers"
	user2 "awesomeProject2/internal/user"
	"awesomeProject2/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	SentPassToEmailTemplate = fmt.Sprint("Ваш пароль %s, используйте его для входа в приложение")
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

	user.Role = "user"
	password := helpers.RandStringPassword(15)
	if err != nil {
		response.Code = http.StatusBadRequest
		log.Println("Не получилось генерировать пароль")
		return
	}

	user.Password, err = helpers.HashPassword(password)
	if err != nil {
		log.Println("HashPassword: пароль не хешировался")
		return
	}

	err = user.CreateUser()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}
	log.Println(user.Password)
	log.Println(password)
	msg := []byte(fmt.Sprintf(SentPassToEmailTemplate, password))
	emails := []string{user.Email}
	err = helpers.SendMessageByEmail(emails, "Ваш пароль", msg)
	if err != nil {
		response.Code = http.StatusBadRequest
		log.Println("Не получилось отправить пароль на почту", err)
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
