package main

import (
	"errors"
	"sync"
)

type Account struct {
	accountNumber string
	balance       float64
	mu            sync.Mutex
}

func NewAccount(number string, initialBalance float64) *Account {
	return &Account{
		accountNumber: number,
		balance:       initialBalance,
	}
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func (a *Account) GetAccountNumber() string {
	return a.accountNumber
}

func (a *Account) Debit(balance float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if balance > a.balance {
		return errors.New("insufficient balance")
	}
	a.balance -= balance
	return nil
}

func (a *Account) Credit(balance float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += balance
	return nil
}
