package main

type LogAppender interface {
	Append(message *LogMessage) error
}
