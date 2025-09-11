package main

import (
	"sync"
)

type BankingService struct {
	Accounts map[string]*Account
	mu       sync.Mutex
}

func NewBankingService() *BankingService {
	return &BankingService{
		Accounts: make(map[string]*Account),
	}
}

func (bs *BankingService) CreateAccount(accountNumber string, amount float64) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.Accounts[accountNumber] = NewAccount(accountNumber, amount)
}

func (bs *BankingService) GetAccount(accountNumber string) *Account {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	return bs.Accounts[accountNumber]
}
