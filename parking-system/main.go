package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	fmt.Println("Parking System")

	parkingLot := GetParkingLotInstance()
	parkingLot.Name = "VTP Parking Lot"

	parkingLot.AddFloor(0)
	parkingLot.AddFloor(1)

	for i := 1; i <= 8; i++ {
		wg.Add(1)

		go parkCar(i, parkingLot)
	}

	wg.Wait()

	parkingLot.DisplayAvailability()
	ticket, _ := parkingLot.ParkVehicle(NewBus("bus-1"))
	parkingLot.DisplayAvailability()
	time.Sleep(5 * time.Second)
	err := parkingLot.UnParkVehicle(ticket)
	if err != nil {
		return
	}

	formattedCharge := fmt.Sprintf("%.2f", ticket.CalculateTotCharge())

	fmt.Printf("bill for %s = %s\n", ticket.Vehicle.GetVehicleNumber(), formattedCharge)

}

func parkCar(ind int, parkingLot *ParkingLot) {
	defer wg.Done()

	car := NewCar(fmt.Sprintf("car-%d", ind))

	ticket, err := parkingLot.ParkVehicle(car)
	if err != nil {
		fmt.Errorf("failed to park %s: %v", car.VehicleNumber, err)
		return
	}

	fmt.Printf("%s parked successfully. Ticket: %s\n", car.VehicleNumber, ticket.ExitTime)

}
