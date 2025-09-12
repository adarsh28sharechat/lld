package main

type Message struct {
	content string
}

func NewMessage(content string) *Message {
	return &Message{
		content: content,
	}
}
