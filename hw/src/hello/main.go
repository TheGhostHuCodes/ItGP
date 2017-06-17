package main

import (
	"fmt"
	"os"
)

func printer(message string) error {
	if message == "" {
		return fmt.Errorf("Unwilling to print an empty string")
	}
	_, err := fmt.Printf("%s\n", message)
	return err
}

func main() {
	if err := printer(""); err != nil {
		fmt.Printf("printer failed: %s\n", err)
		os.Exit(1)
	}
}
