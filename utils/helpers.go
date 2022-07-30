package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type GenericResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}

func WriteResponse(
	w http.ResponseWriter,
	data interface{},
	message string,
	statusCode int,
) {
	resp := GenericResponse{
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}
