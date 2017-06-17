package main

import (
	"fmt"
)

func printer(message string) error {
	defer fmt.Printf("\n")
	_, err := fmt.Printf("%s", message)
	return err
}

func main() {
	printer("Hello, World!")
}
