package entities

import (
	"errors"
)

type PhoneNumber string

type PostClientRequest struct {
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Age         int         `json:"age"`
	PhoneNumber PhoneNumber `json:"phone_number"`
}

func (r PostClientRequest) IsValid() error {
	if r.Age < 0 || r.Age > 120 {
		return errors.New("idade deve estar compreendida entre 0 e 120 anos")
	}

	return nil
}
