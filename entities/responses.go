package entities

import uuid "github.com/satori/go.uuid"

type PostClientResponse struct {
	ID          int         `json:"id"`
	UUID        uuid.UUID   `json:"uuid"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	FullName    string      `json:"full_name"`
	Age         int         `json:"age"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

type GetClientResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
}

type GetClientsResponse struct {
	Total   int                 `json:"total"`
	Clients []GetClientResponse `json:"clients"`
}

func NewPostClientResponse(client Client) PostClientResponse {
	return PostClientResponse{
		ID:          client.ID,
		UUID:        client.UUID,
		FirstName:   client.FirstName,
		LastName:    client.LastName,
		FullName:    client.GetFullName(),
		Age:         client.Age,
		PhoneNumber: client.PhoneNumber,
	}
}

func NewGetClientResponseFrom(client Client) GetClientResponse {
	return GetClientResponse{
		UUID:      client.UUID,
		ID:        client.ID,
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Age:       client.Age,
	}
}

func NewClientsResponseFromClients(clients []Client) GetClientsResponse {
	var clientResponse GetClientsResponse

	for _, client := range clients {
		response := GetClientResponse{
			ID:        client.ID,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Age:       client.Age,
		}

		clientResponse.Clients = append(clientResponse.Clients, response)
	}

	clientResponse.Total = len(clientResponse.Clients)

	return clientResponse
}
