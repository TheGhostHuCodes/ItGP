package main

import (
	"fmt"
	"shuffler"
)

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type weightString struct {
	Weight int
	s      string
}

type stringSlice []weightString

func (ss stringSlice) Len() int {
	return len(ss)
}
func (ss stringSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
func (ss stringSlice) Weight(i int) int {
	return ss[i].Weight
}

func main() {
	i := intSlice{1, 2, 3, 4, 5, 6}
	j := stringSlice{weightString{100, "Hello"}, weightString{200, "World!"},
		weightString{10, "Goodbye"}}

	shuffler.Shuffle(i)
	fmt.Printf("%v\n", i)

	shuffler.WeightedShuffle(j)
	fmt.Printf("%v\n", j)

	fmt.Printf("Loop: %d\n", shuffler.Counter())
}
