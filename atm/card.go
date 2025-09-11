package main

type Card struct {
	CardNumber string
	Pin        string
}

func NewCard(cardNumber, pin string) *Card {
	return &Card{
		CardNumber: cardNumber,
		Pin:        pin,
	}
}
