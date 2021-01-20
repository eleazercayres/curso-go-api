package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/filipefernandes007/curso-go-api/api/entities"
	"github.com/filipefernandes007/curso-go-api/api/services"
	"io/ioutil"
	"log"
	"net/http"
)

func ListClientsHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, fmt.Sprintf("Method %s not expected", request.Method), http.StatusMethodNotAllowed)
		return
	}

	response, err := json.Marshal(services.Clients)

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(response)
	if err != nil {
		http.Error(writer, "error sending response", http.StatusInternalServerError)
	}
}

func CreateClientHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, fmt.Sprintf("Method %s not expected", request.Method), http.StatusMethodNotAllowed)
		return
	}

	// Ler o body/payload do request:
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(writer, "could not read body", http.StatusBadRequest)
		return
	}

	requestClient := entities.PostClientRequest{}
	err = json.Unmarshal(body, &requestClient)
	if err != nil {
		http.Error(writer, "cant't unmarshal body", http.StatusInternalServerError)
		return
	}

	err = requestClient.IsValid()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	service := services.NewClientService()
	client, err := service.CreateClientFrom(requestClient)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	clientResponse := entities.NewPostClientResponse(client)
	response, err := json.Marshal(clientResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	//_, err = fmt.Fprint(writer, string(response))
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write(response)
	if err != nil {
		http.Error(writer, "error sending response", http.StatusInternalServerError)
	}

}
