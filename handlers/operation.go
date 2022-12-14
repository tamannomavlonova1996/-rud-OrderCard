package handlers

import (
	"awesomeProject2/internal/service/operation"
	"awesomeProject2/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateOperation(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.Operation
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = "Не олучилось запарсить данные"
		log.Println(err)
		return
	}
	err = operation.CreateOperation(req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}

	response.Message = "Данные добавились успешно!"
}

func GetOperations(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	operations, err := operation.GetOperations()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		log.Println(err)
		return
	}

	response.Message = "Данные получены успешно!"
	response.Payload = operations
}

func GetOperationByID(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := operation.GetOperationByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные получены успешно!"
	response.Payload = result
}

func UpdateOperationByID(w http.ResponseWriter, r *http.Request) {
	var (
		req      models.Operation
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
	err = operation.UpdateOperationByID(req)

	response.Message = "Данные обновлены успешно!"
}

func DeleteOperationByID(w http.ResponseWriter, r *http.Request) {
	var (
		response = models.Response{
			Code: http.StatusOK,
		}
	)
	defer response.Send(w, r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := operation.DeleteOperationByID(id)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		return
	}
	response.Message = "Данные удалены успешно!"
}
