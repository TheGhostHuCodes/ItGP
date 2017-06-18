package main

import (
	"fmt"
)

func whatIsThis(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("It's a string %s\n", v)
	case uint32:
		fmt.Printf("It's an unsigned 32-bit integer %d\n", v)
	default:
		fmt.Printf("Don't know %v\n", v)
	}
}

func main() {
	whatIsThis(uint32(42))
}
