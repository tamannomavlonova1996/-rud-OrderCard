package handlers

import (
	"awesomeProject2/internal/service/card"
	"awesomeProject2/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.Card
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "Не удалось запарсить данные"
		log.Println(err)
		return
	}
	err = card.CreateCard(req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}

	response.Message = "Данные добавились успешно!"
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	cards, err := card.GetCards()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return

	}

	response.Message = "Данные получены успешно!"
	response.Payload = cards
}

func GetCardByID(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := card.GetCardByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные получены успешно!"
	response.Payload = result
}

func UpdateCardByID(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.Card
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "Не получилось запарсить данные"
		log.Println(err)
		return
	}
	err = card.UpdateCardByID(req)

	response.Message = "Данные обновлены успешно!"
}

func DeleteCardByID(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := card.DeleteCardByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные удалены успешно!"
}
