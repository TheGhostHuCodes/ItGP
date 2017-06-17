package main

import (
	"fmt"
)

func printer(words []string) {
	for _, word := range words {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main() {
	words := make([]string, 4)
	words[0] = "The"
	words[1] = "Quick"
	words[2] = "Brown"
	words[3] = "Fox"
	printer(words)
	fmt.Printf("%d %d\n", len(words), cap(words))
}
