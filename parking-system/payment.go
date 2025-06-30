package main

import "fmt"

type Status string

const (
	PaymentStatusPending   Status = "Pending"
	PaymentStatusCompleted Status = "Completed"
	PaymentStatusFailed    Status = "Failed"
)

type Payment struct {
	Status Status
	Amount float64
	Ticket *ParkingTicket
}

func NewPayment(amount float64, ticket *ParkingTicket) *Payment {
	return &Payment{Status: PaymentStatusPending, Amount: amount, Ticket: ticket}
}

func (p *Payment) GetPaymentStatus() Status {
	return p.Status
}

func (p *Payment) ProcessPayment() error {
	if p.Ticket == nil {
		return fmt.Errorf("ticket not found")
	}
	if p.Ticket.TotoalCharge < p.Amount {
		p.Status = PaymentStatusFailed
		return fmt.Errorf("payment failed: insufficient funds")
	}

	p.Status = PaymentStatusCompleted
	return nil
}
