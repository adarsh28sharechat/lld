package main

import "fmt"

func main() {
	fmt.Println("vending machine")
	machine := NewCoffeeMachine()

	//add coffees
	machine.AddCoffee(NewCoffee("aaa", 4, 20))
	machine.AddCoffee(NewCoffee("bbb", 6, 30))

	machine.DisplayCoffees()
	if err := machine.DispenseCoffee("aaa", 2); err != nil {
		fmt.Errorf("error occurred")
	}

	machine.DisplayCoffees()
}
