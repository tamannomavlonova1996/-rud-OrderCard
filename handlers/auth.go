package handlers

import (
	"awesomeProject2/helpers"
	"awesomeProject2/internal/service"
	user2 "awesomeProject2/internal/user"
	"awesomeProject2/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
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

func SignIn(w http.ResponseWriter, r *http.Request) {
	var (
		request  models.UserLogin
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Code = http.StatusBadRequest
		log.Println(err)
		return
	}

	token, err := service.CreateToken(request.Email, request.Password)
	if err != nil {
		response.Code = http.StatusBadRequest
		log.Println(err)
		return
	}
	response.Payload = token
	response.Message = "Вы успешно зашли в свой аккаунт"
}
