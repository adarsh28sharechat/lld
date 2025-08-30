package main

type Player struct {
	Name   string
	Symbol string
}

func NewPlayer(name string, symbol string) *Player {
	return &Player{Name: name, Symbol: symbol}
}
