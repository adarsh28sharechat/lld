package main

import "fmt"

func main() {
	fmt.Println("Elevator System")

	building := NewBuilding()
	manager := NewElevatorManager(building)

}
