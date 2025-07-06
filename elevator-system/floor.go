package main

type Floor struct {
	FloorNumber int
	FloorPanel  []*HallPanel
}

func NewFloor(floorNumber int) *Floor {
	floor := &Floor{FloorNumber: floorNumber, FloorPanel: make([]*HallPanel, 0)}
	for i := 0; i < 3; i++ {
		floorPanel := NewHallPanel(i, floorNumber)
		floor.FloorPanel = append(floor.FloorPanel, floorPanel)
	}
	return floor
}
