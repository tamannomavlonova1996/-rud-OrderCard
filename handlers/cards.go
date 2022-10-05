package handlers

import (
	"awesomeProject2/internal/cards"
	"awesomeProject2/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateCards(w http.ResponseWriter, r *http.Request) {
	var (
		card     cards.Cards
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "Неверные данные"
		log.Println(err)
		return
	}
	err = card.CreateCards()
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
		card     cards.Cards
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	result, err := card.GetCards()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return

	}

	response.Message = "Данные получены успешно!"
	response.Payload = result
}

func GetCardById(w http.ResponseWriter, r *http.Request) {
	var (
		card     cards.Cards
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

func UpdateCardById(w http.ResponseWriter, r *http.Request) {
	var (
		card     cards.Cards
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "Неверные данные"
		log.Println(err)
		return
	}
	err = card.UpdateCardByID()

	response.Message = "Данные обновлены успешно!"
}

func DeleteCardById(w http.ResponseWriter, r *http.Request) {
	var (
		card     cards.Cards
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
