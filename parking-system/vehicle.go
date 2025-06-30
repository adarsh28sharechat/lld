package main

type VehicleType string

const (
	CarType  VehicleType = "car"
	BikeType VehicleType = "bike"
	BusType  VehicleType = "bus"
)

var vehicleCosts = map[VehicleType]float64{
	CarType:  100,
	BikeType: 50,
	BusType:  200,
}

type Vehicle struct {
	VehicleType   VehicleType `json:"vehicle_type"`
	Cost          float64     `json:"cost"`
	VehicleNumber string      `json:"vehicle_number"`
}

type VehicleInterface interface {
	GetVehicleType() VehicleType
	GetVehicleNumber() string
	GetVehicleCost() float64
}

func (vt Vehicle) GetVehicleType() VehicleType {
	return vt.VehicleType
}

func (vt Vehicle) GetVehicleNumber() string {
	return vt.VehicleNumber
}

func (vt Vehicle) GetVehicleCost() float64 {
	return vt.Cost
}

func NewVehicle(vehicleType VehicleType, vehicleNumber string) *Vehicle {
	cost := vehicleCosts[vehicleType]
	return &Vehicle{VehicleNumber: vehicleNumber, VehicleType: vehicleType, Cost: cost}
}

type Car struct {
	Vehicle
}

func NewCar(vehicleNumber string) *Car {
	return &Car{Vehicle: *NewVehicle(CarType, vehicleNumber)}
}

type Bike struct {
	Vehicle
}

func NewBike(vehicleNumber string) *Bike {
	return &Bike{Vehicle: *NewVehicle(BikeType, vehicleNumber)}
}

type Bus struct {
	Vehicle
}

func NewBus(vehicleNumber string) *Bus {
	return &Bus{Vehicle: *NewVehicle(BusType, vehicleNumber)}
}
