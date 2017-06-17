package main

import (
	"fmt"
	"os"
)

func main() {
	if num_bytes, err := fmt.Printf("Hello, World!\n"); err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d characters\n", num_bytes)
	}
}
