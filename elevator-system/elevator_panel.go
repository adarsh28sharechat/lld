package main

type ElevatorPanel struct {
	PanelID      int
	PanelButtons [15]bool
}

func NewElevatorPanel(panelID int) *ElevatorPanel {
	return &ElevatorPanel{PanelID: panelID, PanelButtons: [15]bool{}}
}
