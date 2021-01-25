package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/filipefernandes007/go-curso/api/entities"
	"github.com/filipefernandes007/go-curso/api/services"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
)

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

func ListClientHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, fmt.Sprintf("Method %s not expected", request.Method), http.StatusMethodNotAllowed)
		return
	}

	service := services.NewClientService()
	clients, err := service.ListClients()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	clientsResponse := entities.NewClientsResponseFromClients(clients)
	response, err := json.Marshal(clientsResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	_, err = writer.Write(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func GetClientHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, fmt.Sprintf("Method %s not expected", request.Method), http.StatusMethodNotAllowed)
		return
	}

	// /api/get-client?uuid=
	var uuidParam = request.URL.Query().Get("uuid")

	// TODO: complete this method

	_uuid, err := uuid.FromString(uuidParam)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	s := services.NewClientService()
	client, err := s.GetClientByUUID(_uuid)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	clientResponse := entities.NewGetClientResponseFrom(client)
	response, err := json.Marshal(&clientResponse)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = writer.Write(response)
	if err != nil {
		http.Error(writer, "error sending response", http.StatusInternalServerError)
		return
	}

	// 1. Obter através do serviço de 'client' (proposta: GetClientByUUID) a entidade.
	// 2. Avaliar se há erro, em que o erro será espoletado caso não tenha encontrado nenhuma entidade.
	// 3. Digerir o erro, transformando-o num 404.
	// 4. Transformar a entidade 'client', numa resposta apropriada: GetClientResponse
	// 5. Marshal do 'client response' : ver exemplos anteriores.
	// 6. 'Reply' do handler com base na resposta anterior: ver exemplos anteriores.
	// 7. Testar:
	// 		7.1 - Criar uma entidade: curl -v -X POST -d '{"first_name":"Sharon","last_name":"Stone","age":62,"phone_number":"999999999"}
	//		7.2 - Testar novo endpoint: curl http://localhost:9000/api/get-client\?uuid\=<uuid obtido na criação>


}
