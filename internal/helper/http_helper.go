package helper

import (
	"fmt"
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error, statuscode int) {
	// Error function will write the string into the w  means ResponseWriter
	// Sprintf will return a string after placeing err message in place of %v with statuscode
	http.Error(w, fmt.Sprintf("Error: %v\n", err), statuscode)
	log.Printf("Error: %v", err)
}

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
