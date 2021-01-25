package services

import (
	"github.com/filipefernandes007/curso-go-api/api/entities"
	uuid "github.com/satori/go.uuid"
	"math/rand"
)

var Clients []entities.Client = make([]entities.Client, 0)

type ClientService struct {
}

func NewClientService() ClientService {
	return ClientService{}
}

func (s ClientService) CreateClientFrom(request entities.PostClientRequest) (entities.Client, error) {
	client := entities.Client{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Age:         request.Age,
		PhoneNumber: request.PhoneNumber,
	}

	client.ID = rand.Intn(1000)
	client.UUID = uuid.NewV4()

	Clients = append(Clients, client)

	return client, nil
}
