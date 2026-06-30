package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func WriteJson(w http.ResponseWriter, status int, data any) {
	out, err := json.Marshal(data)
	if err != nil {
		log.Println("unable to marshal json:", err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		log.Println("unable to write json:", err)
	}
}

func ErrJson(w http.ResponseWriter, err error) {

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	status, exist := CustomError[err]
	if exist {
		WriteJson(w, status, payload)
	}

	WriteJson(w, http.StatusInternalServerError, payload)
}
