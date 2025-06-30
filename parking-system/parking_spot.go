package main

import (
	"fmt"
	"sync"
)

type ParkingSpot struct {
	SpotId         int
	VehicleType    VehicleType
	CurrentVehicle *VehicleInterface
	lock           sync.Mutex
}

func NewParkingSpot(parkingSpotId int, parkingSpotType VehicleType) *ParkingSpot {
	return &ParkingSpot{SpotId: parkingSpotId, VehicleType: parkingSpotType}
}

func (p *ParkingSpot) IsParkingSpotFree() bool {
	return p.CurrentVehicle == nil
}

func (p *ParkingSpot) RemoveVehile() {
	p.CurrentVehicle = nil
}

func (p *ParkingSpot) ParkVehicle(vehicle VehicleInterface) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if vehicle.GetVehicleType() != p.VehicleType {
		return fmt.Errorf("vehicle type mismatch: expected %s, got %s", p.VehicleType, vehicle.GetVehicleType())
	}

	if p.CurrentVehicle != nil {
		return fmt.Errorf("parking spot %s is already parked", p.SpotId)
	}

	p.CurrentVehicle = &vehicle
	return nil
}
