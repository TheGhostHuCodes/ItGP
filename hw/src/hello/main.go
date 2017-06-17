package main

import (
	"fmt"
	"time"
)

func emit(c chan string, done chan bool) {
	defer close(c)
	words := []string{"The", "quick", "brown", "fox"}
	timer := time.NewTimer(3 * time.Second)
	i := 0
	for {
		select {
		case c <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-done:
			done <- true
			return
		case <-timer.C:
			return
		}
	}
}

func main() {
	wordChannel := make(chan string)
	doneChannel := make(chan bool)

	go emit(wordChannel, doneChannel)

	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
}
