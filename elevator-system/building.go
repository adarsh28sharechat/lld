package main

type Building struct {
	Floors    []*Floor
	Elevators []*Elevator
}

func NewBuilding() *Building {
	building := &Building{Floors: make([]*Floor, 0)}

	for i := 1; i <= 15; i++ {
		newFloor := NewFloor(i)
		building.Floors = append(building.Floors, newFloor)
	}

	for i := 1; i <= 3; i++ {
		newElevator := NewElevator(i)
		building.Elevators = append(building.Elevators, newElevator)
	}

	return building
}
