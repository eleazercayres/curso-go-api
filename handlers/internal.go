package handlers

import (
	"log"
	"net/http"
)

func HealthyHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("I'm alive!"))
	if err != nil {
		log.Fatal("Not healthy")
	}
}
