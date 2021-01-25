package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)

type Address struct {
	FirstAddress string
}

type Client struct {
	ID          int
	UUID        uuid.UUID
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber PhoneNumber
	Address     Address
}

func (c *Client) SetFirstName(n string) error {
	if n == "" {
		return errors.New("empty value")
	}

	c.FirstName = n

	return nil
}

func (c Client) GetFullName() string {
	return c.FirstName + " " + c.LastName
}
