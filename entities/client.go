package entities

import uuid "github.com/satori/go.uuid"

type Client struct {
	ID          int
	UUID        uuid.UUID
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber PhoneNumber
}

func (c Client) GetFullName() string {
	return c.FirstName + " " + c.LastName
}
