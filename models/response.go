package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

// Send ...
func (res *Response) Send(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(res.Code)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Sending response failed:", err)
	}

}
