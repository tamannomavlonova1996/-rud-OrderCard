package handlers

import (
	"awesomeProject2/helpers"
	"awesomeProject2/internal/service/user"
	"awesomeProject2/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.User
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

	err = user.CreateUser(req)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		log.Println(err)
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

	token, err := helpers.CreateToken(request.Email, request.Password)
	if err != nil {
		response.Code = http.StatusBadRequest
		log.Println(err)
		return
	}
	response.Payload = token
	response.Message = "Вы успешно зашли в свой аккаунт"
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.ResetPassword
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
	err = user.ResetPassword(&req)
	if err != nil {
		response.Code = http.StatusBadRequest
		log.Println(err)
		return
	}
	response.Message = "Вы успешно забросили свой пароль"

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r.Context())
	fmt.Println(userID)
}
