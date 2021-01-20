package services

import (
	"github.com/filipefernandes007/curso-go-api/api/entities"
	"math/rand"
)

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
	return client, nil
}
