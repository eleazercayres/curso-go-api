package services

import (
	"errors"
	"github.com/filipefernandes007/go-curso/api/entities"
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

func (s ClientService) ListClients() ([]entities.Client, error) {
	return Clients, nil
}

func (s ClientService) GetClientByUUID(uuid uuid.UUID) (entities.Client, error) {
	clients, err := s.ListClients()
	if err != nil {
		return entities.Client{}, err
	}

	for _, client := range clients {
		if client.UUID == uuid {
			return client, nil
		}
	}

	return entities.Client{}, errors.New("client not found")
}

