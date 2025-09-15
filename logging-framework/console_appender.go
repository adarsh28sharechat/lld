package main

import "fmt"

type ConsoleAppender struct{}

func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{}
}

func (ca *ConsoleAppender) Append(message *LogMessage) error {
	fmt.Println(message.String())
	return nil
}
