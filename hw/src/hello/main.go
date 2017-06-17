package main

import (
	"fmt"
)

const (
	message = "The answer to life the universe and everything is %d\n"
	answer  = 42
)

func main() {
	fmt.Printf(message, answer)
}
