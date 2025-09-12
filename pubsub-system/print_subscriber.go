package main

import "fmt"

type PrintSubscriber struct {
	Name string
}

func NewPrintSubscriber(name string) *PrintSubscriber {
	return &PrintSubscriber{
		Name: name,
	}
}

func (ps *PrintSubscriber) OnMessage(message *Message) {
	fmt.Printf("Subscriber %s received message: %s\n", ps.Name, message.content)
}
