package main

type Group struct {
	Id       string
	Name     string
	Members  []*User
	Expenses []*Expense
}

func NewGroup(id string, name string) *Group {
	return &Group{Id: id, Name: name, Members: []*User{}, Expenses: []*Expense{}}
}

func (g *Group) AddMember(user *User) {
	g.Members = append(g.Members, user)
}
