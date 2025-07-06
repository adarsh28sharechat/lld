package main

import "fmt"

type Directions string

const (
	Up    Directions = "Up"
	Down  Directions = "Down"
	Still Directions = "Still"
)

type HallPanel struct {
	PanelID              int
	DirectionInstruction Directions
	SourceFloor          int
}

func NewHallPanel(panelID int, sourceFloor int) *HallPanel {
	return &HallPanel{PanelID: panelID, DirectionInstruction: Still, SourceFloor: sourceFloor}
}

func (h *HallPanel) AddDirectionInstruction(directionInstruction Directions) {
	h.DirectionInstruction = directionInstruction
}

func (h *HallPanel) RequestElevator(manager *ElevatorManager, directions Directions) (elevator *Elevator) {
	fmt.Printf("panel %d requesting elevator from floor %d with directions %s\n", h.PanelID, h.SourceFloor, directions)
	return manager.
}
