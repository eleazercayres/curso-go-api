package entities

type Client struct {
	ID          int
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber PhoneNumber
}

func (c Client) GetFullName() string {
	return c.FirstName + " " + c.LastName
}
