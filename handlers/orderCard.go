package handlers

import (
	"awesomeProject2/internal/orderCard"
	"awesomeProject2/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateOrderCard(w http.ResponseWriter, r *http.Request) {
	var (
		card     orderCard.OrderCard
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
	err = card.CreateOrderCards()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}

	response.Message = "Данные добавились успешно!"
}

func GetOrderCards(w http.ResponseWriter, r *http.Request) {
	var (
		card     orderCard.OrderCard
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	result, err := card.GetOrderCards()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return

	}

	response.Message = "Данные получены успешно!"
	response.Payload = result
}

func GetOrderCardByID(w http.ResponseWriter, r *http.Request) {
	var (
		card     orderCard.OrderCard
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := card.GetOrderCardByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные получены успешно!"
	response.Payload = result
}

func UpdateOrderCardByID(w http.ResponseWriter, r *http.Request) {
	var (
		card     orderCard.OrderCard
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
	err = card.UpdateOrderCardByID()

	response.Message = "Данные обновлены успешно!"
}

func DeleteOrderCardByID(w http.ResponseWriter, r *http.Request) {
	var (
		card     orderCard.OrderCard
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := card.DeleteOrderCardByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные удалены успешно!"
}
