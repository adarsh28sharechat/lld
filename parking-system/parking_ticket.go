package main

import "time"

const baseCharge = 100.00

type ParkingTicket struct {
	EntryTime    time.Time
	ExitTime     time.Time
	Vehicle      VehicleInterface
	Spot         *ParkingSpot
	TotoalCharge float64
}

func NewParkingTicket(vehicle VehicleInterface, spot *ParkingSpot) *ParkingTicket {
	return &ParkingTicket{EntryTime: time.Now(), ExitTime: time.Time{}, Vehicle: vehicle, Spot: spot, TotoalCharge: 0.00}
}

func (p *ParkingTicket) SetExitTime(exitTime time.Time) {
	p.ExitTime = exitTime
}

func (p *ParkingTicket) CalculateTotCharge() float64 {
	if p.ExitTime == (time.Time{}) {
		p.TotoalCharge = baseCharge
		return p.TotoalCharge
	}
	duration := p.EntryTime.Sub(p.EntryTime)
	hours := duration.Hours()
	additionalCost := hours * p.Vehicle.GetVehicleCost()
	p.TotoalCharge = baseCharge + additionalCost
	return p.TotoalCharge
}
