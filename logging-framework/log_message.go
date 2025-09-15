package main

import (
	"fmt"
	"time"
)

type LogMessage struct {
	Message   string
	Timestamp int64
	Level     LogLevel
}

func NewLogMessage(level LogLevel, message string) *LogMessage {
	return &LogMessage{
		Message:   message,
		Level:     level,
		Timestamp: time.Now().UnixMilli(),
	}
}

func (lm *LogMessage) String() string {
	return fmt.Sprintf("[%s] %d - %s", lm.Level, lm.Timestamp, lm.Message)
}
