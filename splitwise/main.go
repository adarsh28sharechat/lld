package main

import "fmt"

func main() {
	fmt.Println("Splitwise")

	service := GetSplitwiseService()

	user1 := NewUser("1", "adarsh")
	user2 := NewUser("2", "sahil")
	user3 := NewUser("3", "paajesh")

	service.AddUser(user1)
	service.AddUser(user2)
	service.AddUser(user3)

	group1 := NewGroup("1", "A006")
	group1.AddMember(user1)
	group1.AddMember(user2)
	group1.AddMember(user3)

	service.AddGroup(group1)

	expense1 := NewExpense("1", "Rent", 300, user1, group1)
	service.DistributeExpenseInGroup(expense1)

	service.ShowGroupUsersBalance(group1)

	expense2 := NewExpense("2", "cook", 400, user2, group1)
	service.DistributeExpenseInGroup(expense2)

	service.ShowGroupUsersBalance(group1)
}

//user can add other users and can create groups
//Add expense divide equally among participants
//suggest payments
