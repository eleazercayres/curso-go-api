package entities

type Address struct {
	FirstAddress string
}

type Client struct {
	ID          int
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber PhoneNumber
	Address     Address
}

func (c Client) GetFullName() string {
	return c.FirstName + " " + c.LastName
}
