package main

import (
	"fmt"
	"poetry"
)

func main() {
	p, err := poetry.LoadPoem("wordsworth.txt")
	if err != nil {
		fmt.Printf("Error loading poem: %s\n", err)
	}
	v, c, puncs := p.Stats()
	fmt.Printf("Vowels: %d, Consonants: %d, Punctuation: %d\n", v, c, puncs)
	fmt.Printf("Stanzas: %d, Lines: %d\n", p.NumStanzas(), p.NumLines())
	fmt.Printf("%s", p)
}
