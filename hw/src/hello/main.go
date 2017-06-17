package main

import (
	"fmt"
)

func emit(c chan string, done chan bool) {
	words := []string{"The", "quick", "brown", "fox"}
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
		}
	}
}

func main() {
	wordChannel := make(chan string)
	doneChannel := make(chan bool)

	go emit(wordChannel, doneChannel)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordChannel)
	}
	doneChannel <- true
	<-doneChannel
}
