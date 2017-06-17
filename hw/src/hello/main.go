package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}
}
