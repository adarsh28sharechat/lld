package main

import (
	"fmt"
	"sync"
)

type CashDispenser struct {
	AvailableCash float64
	mu            sync.Mutex
}

func NewCashDispenser(initialCash float64) *CashDispenser {
	return &CashDispenser{
		AvailableCash: initialCash,
	}
}

func (d *CashDispenser) DispenseCash(amount float64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if amount > d.AvailableCash {
		return fmt.Errorf("cannot dispense Cash %f", amount)
	}
	d.AvailableCash -= amount
	return nil
}
