package handlers

import (
	"awesomeProject2/internal/service/orderCard"
	"awesomeProject2/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateOrderCard(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.OrderCard
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "Неверные данные"
		log.Println(err)
		return
	}
	err = orderCard.CreateOrderCard(req)
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
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	result, err := orderCard.GetOrderCards()
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
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := orderCard.GetOrderCardByID(id)
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
		req      models.OrderCard
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
	err = orderCard.UpdateOrderCardByID(req)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "Не удалось обновить данные "
		log.Println(err)
		return
	}

	response.Message = "Данные обновлены успешно!"
}

func DeleteOrderCardByID(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := orderCard.DeleteOrderCardByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные удалены успешно!"
}
