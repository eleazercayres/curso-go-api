package entities

type PostClientResponse struct {
	ID          int         `json:"id"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	FullName    string      `json:"full_name"`
	Age         int         `json:"age"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

func NewPostClientResponse(client Client) PostClientResponse {
	return PostClientResponse{
		ID:          client.ID,
		FirstName:   client.FirstName,
		LastName:    client.LastName,
		FullName:    client.GetFullName(),
		Age:         client.Age,
		PhoneNumber: client.PhoneNumber,
	}
}
