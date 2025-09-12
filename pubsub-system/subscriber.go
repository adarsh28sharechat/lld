package main

type Subscriber interface {
	OnMessage(message *Message)
}
