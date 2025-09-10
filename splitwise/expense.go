package main

type Expense struct {
	Id          string
	Description string
	Amount      float64
	PaidBy      *User
	Group       *Group
}

func NewExpense(id string, desc string, amount float64, payer *User, group *Group) *Expense {
	return &Expense{
		Id:          id,
		Description: desc,
		Amount:      amount,
		PaidBy:      payer,
		Group:       group,
	}
}
