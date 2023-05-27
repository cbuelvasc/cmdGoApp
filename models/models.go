package models

type Client struct {
	Id    int
	Name  string
	Email string
	Phone string
}

type Clients []Client
