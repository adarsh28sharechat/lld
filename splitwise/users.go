package main

type User struct {
	Id     string
	Name   string
	Amount float64
}

func NewUser(id string, name string) *User {
	return &User{Id: id, Name: name, Amount: 0}
}
