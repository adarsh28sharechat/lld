package main

import "sync"

type Logger struct {
	Config *LoggerConfig
	mu     sync.Mutex
}

var instance *Logger
var once sync.Once

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			Config: NewLoggerConfig(LogLevelInfo, NewConsoleAppender()),
		}
	})
	return instance
}

func (l *Logger) SetConfig(config *LoggerConfig) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Config = config
}

func (l *Logger) log(level LogLevel, message string) error {
	l.mu.Lock()
	if level < l.Config.Level {
		defer l.mu.Unlock()
		return nil
	}
	appender := l.Config.Appender
	defer l.mu.Unlock()

	logMessage := NewLogMessage(level, message)
	return appender.Append(logMessage)
}

func (l *Logger) Info(message string) error {
	return l.log(LogLevelInfo, message)
}

func (l *Logger) Error(message string) error {
	return l.log(LogLevelError, message)
}

func (l *Logger) Debug(message string) error {
	return l.log(LogLevelDebug, message)
}

func (l *Logger) Fatal(message string) error {
	return l.log(LogLevelFatal, message)
}

func (l *Logger) Warning(message string) error {
	return l.log(LogLevelWarning, message)
}
