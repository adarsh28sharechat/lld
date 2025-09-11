package main

import "errors"

type ATM struct {
	BankingService *BankingService
	CashDispenser  *CashDispenser
}

func NewATM(service *BankingService, dispenser *CashDispenser) *ATM {
	return &ATM{
		BankingService: service,
		CashDispenser:  dispenser,
	}
}

func (atm *ATM) AuthenticateUser(card *Card) error {
	return nil
}

func (atm *ATM) CheckBalance(accountNumber string) (float64, error) {
	account := atm.BankingService.GetAccount(accountNumber)
	if account == nil {
		return 0, errors.New("account not found")
	}
	return account.GetBalance(), nil
}

func (atm *ATM) WithdrawalCash(accountNumber string, amount float64) (float64, error) {
	account := atm.BankingService.GetAccount(accountNumber)
	accountBalance := account.GetBalance()
	if accountBalance < amount {
		return 0, errors.New("insufficient funds")
	}
	account.balance -= amount
	return account.balance, nil
}
