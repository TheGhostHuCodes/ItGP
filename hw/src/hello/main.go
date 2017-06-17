package main

import (
	"fmt"
)

func printer(messages ...string) {
	for _, message := range messages {
		fmt.Printf("%s\n", message)
	}
}

func main() {
	printer("Hello,", "World!")
}
