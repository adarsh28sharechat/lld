package main

import (
	"fmt"
	"sync"
)

type CoffeeMachine struct {
	coffees []*Coffee
	mutex   *sync.Mutex
}

var instance *CoffeeMachine
var once sync.Once

func NewCoffeeMachine() *CoffeeMachine {
	once.Do(func() {
		instance = &CoffeeMachine{
			coffees: make([]*Coffee, 0),
			mutex:   &sync.Mutex{},
		}
	})
	return instance
}

func (c *CoffeeMachine) AddCoffee(coffee *Coffee) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.coffees = append(c.coffees, coffee)
}

func (c *CoffeeMachine) DispenseCoffee(coffeeName string, quantity int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for _, coffee := range c.coffees {
		if coffee.name == coffeeName {
			if coffee.quantity < quantity {
				return fmt.Errorf("coffee not available")
			}
			coffee.quantity -= quantity
		}
	}
	return nil
}

func (c *CoffeeMachine) DisplayCoffees() {
	for _, coffee := range c.coffees {
		if coffee.quantity > 0 {
			fmt.Printf("available coffee: %s, in quantity %d \n", coffee.name, coffee.quantity)
		}
	}
}
