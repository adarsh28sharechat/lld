package main

import "fmt"

func main() {
	fmt.Println("ATM")
	service := NewBankingService()
	cashDispenser := NewCashDispenser(10000)

	atm := NewATM(service, cashDispenser)

	service.CreateAccount("12345", 2000)
	service.CreateAccount("12346", 3000)

	//create card and authenticate
	card := NewCard("123456789", "1234")
	if err := atm.AuthenticateUser(card); err != nil {
		fmt.Println("user not authenticated")
		return
	}

	//check balance
	if balance, err := atm.CheckBalance("12345"); err == nil {
		fmt.Println("balance:", balance)
	} else {
		fmt.Println("error:", err)
	}

	//withdrawal cash
	if amount, err := atm.WithdrawalCash("12345", 500); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("remaining cash: %f\n", amount)
	}
}
