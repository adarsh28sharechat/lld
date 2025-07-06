package main

import "sync"

type Elevator struct {
	ElevatorID       int
	Destinations     []int
	CurrentFloor     int
	Capacity         int
	CurrentLoad      int
	Elevator         *ElevatorPanel
	CurrentDirection Directions
	sync.Mutex
}

func NewElevator(id int) *Elevator {
	return &Elevator{ElevatorID: id, Capacity: 10, CurrentLoad: 0, Elevator: NewElevatorPanel(id), CurrentFloor: 1, CurrentDirection: Still}
}

func (e *Elevator) UpdateCurrentFloor(newFloor int) {
	e.Lock()
	e.CurrentFloor = newFloor
	e.Unlock()
}

func (e *Elevator) UpdateCapacity(newCapacity int) {
	e.Lock()
	e.Capacity = newCapacity
	e.Unlock()
}

func (e *Elevator) UpdateCurrentLaod(newLoad int) {
	e.Lock()
	e.CurrentLoad = newLoad
	e.Unlock()
}

func (e *Elevator) UpdateCurrentDirection(newDirection Directions) {
	e.Lock()
	e.CurrentDirection = newDirection
	e.Unlock()
}

func (e *Elevator) FarthestDistance() int {
	maxFloor := 0
	for _, floor := range e.Destinations {
		if floor > maxFloor {
			maxFloor = floor
		}
	}

	return maxFloor
}

func (e *Elevator) NearestDistance() int {
	minFloor := 100

	for _, floor := range e.Destinations {
		if floor < minFloor {
			minFloor = floor
		}
	}
	return minFloor
}
