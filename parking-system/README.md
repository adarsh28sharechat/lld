# Parking System Design

A comprehensive parking lot management system implemented in Go, demonstrating object-oriented design principles and concurrent programming.

## Table of Contents
- [Overview](#overview)
- [System Architecture](#system-architecture)
- [Design Patterns Used](#design-patterns-used)
- [Core Components](#core-components)
- [Step-by-Step Design Process](#step-by-step-design-process)
- [File Structure](#file-structure)
- [How to Run](#how-to-run)
- [Features](#features)
- [Concurrency Handling](#concurrency-handling)
- [Future Enhancements](#future-enhancements)

## Overview

This parking system is designed to manage a multi-floor parking lot that can accommodate different types of vehicles (Cars, Bikes, and Buses). The system handles vehicle parking, unparking, payment processing, and real-time availability tracking with thread-safe operations.

## System Architecture

The system follows a layered architecture with clear separation of concerns:

```
┌─────────────────────────────────────┐
│           Main Application          │
├─────────────────────────────────────┤
│         Parking Lot (Singleton)     │
├─────────────────────────────────────┤
│              Floors                 │
├─────────────────────────────────────┤
│           Parking Spots             │
├─────────────────────────────────────┤
│    Vehicles │ Tickets │ Payments    │
└─────────────────────────────────────┘
```

## Design Patterns Used

### 1. Singleton Pattern
- **Where**: `ParkingLot` struct
- **Why**: Ensures only one parking lot instance exists throughout the application
- **Implementation**: Uses `sync.Once` for thread-safe initialization

### 2. Factory Pattern
- **Where**: Vehicle creation (`NewCar`, `NewBike`, `NewBus`)
- **Why**: Provides a clean interface for creating different vehicle types
- **Implementation**: Factory functions that encapsulate vehicle creation logic

### 3. Strategy Pattern
- **Where**: Vehicle interface and implementations
- **Why**: Allows different vehicle types to have different behaviors while maintaining a common interface

## Core Components

### 1. Vehicle System (`vehicle.go`)

**Purpose**: Defines vehicle types and their properties

**Key Elements**:
- `VehicleInterface`: Common interface for all vehicles
- `VehicleType`: Enum for vehicle categories (Car, Bike, Bus)
- Vehicle cost mapping for billing calculations
- Concrete implementations: `Car`, `Bike`, `Bus`

**Design Decision**: Used interface-based design for polymorphism and extensibility

### 2. Parking Spot (`parking_spot.go`)

**Purpose**: Represents individual parking spaces

**Key Features**:
- Thread-safe operations using `sync.Mutex`
- Vehicle type validation
- Occupancy status tracking
- Atomic park/unpark operations

**Concurrency**: Mutex ensures thread-safe parking operations

### 3. Floor Management (`floor.go`)

**Purpose**: Manages parking spots on each floor

**Key Features**:
- Configurable spot allocation per vehicle type
- Spot availability tracking
- Floor-level search operations

**Configuration**:
```go
const (
    CarSpotCount  = 5
    BikeSpotCount = 10
    BusSpotCount  = 2
)
```

### 4. Parking Lot (`parking_lot.go`)

**Purpose**: Main controller implementing Singleton pattern

**Key Responsibilities**:
- Fleet management across multiple floors
- Vehicle parking/unparking coordination
- System-wide availability display
- Spot allocation strategy

### 5. Ticketing System (`parking_ticket.go`)

**Purpose**: Handles parking tickets and billing

**Key Features**:
- Entry/exit time tracking
- Dynamic charge calculation
- Base charge + hourly rate model

**Billing Formula**:
```
Total Charge = Base Charge + (Hours × Vehicle Rate)
```

### 6. Payment Processing (`payment.go`)

**Purpose**: Manages payment transactions

**Key Features**:
- Payment status tracking
- Transaction validation
- Error handling for failed payments

## Step-by-Step Design Process

### Step 1: Requirements Analysis
We identified the core requirements:
- Multi-floor parking lot
- Support for different vehicle types
- Real-time availability tracking
- Billing and payment system
- Concurrent access support

### Step 2: Entity Identification
Key entities identified:
- **ParkingLot**: Main system controller
- **Floor**: Represents each parking floor
- **ParkingSpot**: Individual parking spaces
- **Vehicle**: Cars, bikes, buses with different properties
- **ParkingTicket**: Tracks parking sessions
- **Payment**: Handles billing transactions

### Step 3: Relationship Modeling
```
ParkingLot (1) ──── (N) Floor
Floor (1) ──── (N) ParkingSpot
ParkingSpot (1) ──── (0..1) Vehicle
Vehicle (1) ──── (1) ParkingTicket
ParkingTicket (1) ──── (1) Payment
```

### Step 4: Interface Design
Defined clear interfaces for:
- `VehicleInterface`: Common vehicle operations
- Consistent method signatures across components

### Step 5: Concurrency Design
Implemented thread safety using:
- `sync.Once` for singleton initialization
- `sync.Mutex` for parking spot operations
- `sync.WaitGroup` for coordinating goroutines

### Step 6: Error Handling Strategy
Comprehensive error handling for:
- Parking spot unavailability
- Vehicle type mismatches
- Payment failures
- Invalid operations

## File Structure

```
parking-system/
├── main.go              # Application entry point and demo
├── parking_lot.go       # Singleton parking lot controller
├── floor.go            # Floor management and spot allocation
├── parking_spot.go     # Individual parking space logic
├── vehicle.go          # Vehicle types and interface
├── parking_ticket.go   # Ticketing and billing system
├── payment.go          # Payment processing
├── go.mod             # Go module definition
└── README.md          # This documentation
```

## How to Run

### Prerequisites
- Go 1.23.5 or later

### Steps
1. Navigate to the parking-system directory:
   ```bash
   cd parking-system
   ```

2. Run the application:
   ```bash
   go run .
   ```

### Expected Output
The application will:
1. Create a parking lot with 2 floors
2. Concurrently park 8 cars using goroutines
3. Display availability status
4. Park a bus and show updated availability
5. Process unparking and payment for the bus

## Features

### Current Features
- ✅ Multi-floor parking management
- ✅ Support for Cars, Bikes, and Buses
- ✅ Real-time availability tracking
- ✅ Concurrent parking operations
- ✅ Dynamic billing system
- ✅ Payment processing
- ✅ Thread-safe operations
- ✅ Comprehensive error handling

### System Capabilities
- **Scalability**: Easy to add new floors and vehicle types
- **Concurrency**: Handles multiple simultaneous parking requests
- **Reliability**: Thread-safe operations prevent race conditions
- **Extensibility**: Interface-based design allows easy feature additions

## Concurrency Handling

### Thread Safety Mechanisms
1. **Singleton Safety**: `sync.Once` ensures single instance creation
2. **Parking Operations**: `sync.Mutex` prevents race conditions
3. **Goroutine Coordination**: `sync.WaitGroup` manages concurrent operations

### Concurrent Scenarios Handled
- Multiple vehicles trying to park simultaneously
- Concurrent access to the same parking spot
- Parallel floor searches
- Simultaneous payment processing

## Future Enhancements

### Potential Improvements
1. **Database Integration**: Persistent storage for tickets and payments
2. **Reservation System**: Advance booking capabilities
3. **Mobile App Integration**: REST API endpoints
4. **Real-time Notifications**: WebSocket-based updates
5. **Advanced Pricing**: Time-based and location-based pricing
6. **Reporting System**: Analytics and usage reports
7. **Security Features**: Access control and surveillance integration
8. **Multi-location Support**: Chain of parking lots

### Technical Enhancements
- Configuration file support
- Logging framework integration
- Health check endpoints
- Metrics collection
- Docker containerization
- CI/CD pipeline setup

## Design Principles Applied

1. **Single Responsibility Principle**: Each struct has a focused responsibility
2. **Open/Closed Principle**: Easy to extend with new vehicle types
3. **Interface Segregation**: Clean, focused interfaces
4. **Dependency Inversion**: Depends on abstractions, not concretions
5. **DRY (Don't Repeat Yourself)**: Reusable components and functions
6. **SOLID Principles**: Overall adherence to SOLID design principles

This parking system demonstrates a well-structured, concurrent, and extensible design that can serve as a foundation for real-world parking management applications. 