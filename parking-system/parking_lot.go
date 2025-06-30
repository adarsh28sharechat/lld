package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	parkingLotInstance *ParkingLot
	once               sync.Once
)

type ParkingLot struct {
	Name   string
	Floors []*Floor
}

func GetParkingLotInstance() *ParkingLot {

	once.Do(func() {
		parkingLotInstance = &ParkingLot{}
	})
	return parkingLotInstance
}

func (p *ParkingLot) AddFloor(floorId int) {
	p.Floors = append(p.Floors, NewParkingFloor(floorId))
}

func (p *ParkingLot) DisplayAvailability() {
	fmt.Printf("parking lots: %s\n", p.Name)

	for _, floor := range p.Floors {
		floor.DisplayFloorStatus(floor)
	}
}

func (p *ParkingLot) findParkingSpot(vehicleType VehicleType) (*ParkingSpot, error) {
	for _, floor := range p.Floors {
		if spot := floor.FindParkingSpot(vehicleType); spot != nil {
			return spot, nil
		}
	}
	return nil, fmt.Errorf("parking spot not found")
}

func (p *ParkingLot) ParkVehicle(vehicle VehicleInterface) (*ParkingTicket, error) {
	parkingSpot, err := p.findParkingSpot(vehicle.GetVehicleType())
	if err != nil {
		return nil, err
	}

	err = parkingSpot.ParkVehicle(vehicle)
	if err != nil {
		return nil, err
	}

	parkingTicket := NewParkingTicket(vehicle, parkingSpot)
	return parkingTicket, nil
}

func (p *ParkingLot) UnParkVehicle(ticket *ParkingTicket) error {
	ticket.SetExitTime(time.Now())
	charge := ticket.CalculateTotCharge()

	payment := NewPayment(charge, ticket)

	if err := payment.ProcessPayment(); err != nil {
		return fmt.Errorf("payment failed: %v. Vehicle is still parked", err)
	}

	ticket.Spot.RemoveVehile()
	return nil
}
