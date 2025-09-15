package main

import "fmt"

func main() {
	fmt.Println("logging-framework")

	logger := GetLogger()

	logger.Info("This is an information message")
	logger.Warning("This is a warning message")
	logger.Error("This is an error message")

	if err := logger.Debug("This is a debug message"); err != nil {
		fmt.Printf("Error logging debug message: %v\n", err)
	}

	if err := logger.Info("This is an information message"); err != nil {
		fmt.Printf("Error logging info message: %v\n", err)
	}
}
