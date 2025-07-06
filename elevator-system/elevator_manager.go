package main

import "math"

type ElevatorManager struct {
	Building *Building
}

func NewElevatorManager(building *Building) *ElevatorManager {
	return &ElevatorManager{Building: building}
}

func (e *ElevatorManager) AssignElevator(floorNumber int, direction Directions) (elevator *Elevator) {
	//bestElevator :=
}

func (e *ElevatorManager) FindClosestElevator(floorNumber int, direction Directions) (elevator *Elevator) {
	var closestElevator *Elevator
	minDistance := int(1e9)
	for _, elevator = range e.Building.Elevators {
		elevator.Lock()

		distance := e.calculateDistance(elevator, floorNumber, direction)
		if distance < minDistance {
			minDistance = distance
			closestElevator = elevator
		}
	}
	return closestElevator
}

func (e *ElevatorManager) calculateDistance(elevator *Elevator, floorNumber int, direction Directions) int {
	currentFloor := elevator.CurrentFloor
	currentDirection := elevator.CurrentDirection

	if currentDirection == Still || currentDirection == direction {
		return currentFloor - floorNumber
	}

	if currentDirection != direction {

	}

}
