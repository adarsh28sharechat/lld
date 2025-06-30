package main

import "fmt"

type FloorType int

const (
	CarSpotCount  = 5
	BikeSpotCount = 10
	BusSpotCount  = 2
)

type Floor struct {
	FloorID      int
	ParkingSpots map[VehicleType]map[int]*ParkingSpot
}

func NewParkingFloor(floorId int) *Floor {
	parkingSpot := make(map[VehicleType]map[int]*ParkingSpot)

	parkingSpot[CarType] = createParkingSpot(CarSpotCount, CarType)
	parkingSpot[BusType] = createParkingSpot(BusSpotCount, BusType)
	parkingSpot[BikeType] = createParkingSpot(BikeSpotCount, BikeType)

	return &Floor{FloorID: floorId, ParkingSpots: parkingSpot}
}

func createParkingSpot(count int, vehicleType VehicleType) map[int]*ParkingSpot {
	spots := make(map[int]*ParkingSpot)
	for i := 0; i < count; i++ {
		spots[i] = NewParkingSpot(i, vehicleType)
	}
	return spots
}

func (f *Floor) DisplayFloorStatus(floor *Floor) {
	fmt.Printf("Floor ID: %d\n", floor.FloorID)

	for vehicleType, parkingSpot := range f.ParkingSpots {
		//fmt.Printf("vehicle type : %s\n", vehicleType)

		count := 0
		for _, parkingSpot := range parkingSpot {
			if parkingSpot.IsParkingSpotFree() {
				count++
			}
		}

		fmt.Printf("spot type: %s, count : %d\n", vehicleType, count)
	}
}

func (f *Floor) FindParkingSpot(vehicleType VehicleType) *ParkingSpot {
	for _, spot := range f.ParkingSpots[vehicleType] {
		if spot.IsParkingSpotFree() {
			return spot
		}
	}
	return nil
}
