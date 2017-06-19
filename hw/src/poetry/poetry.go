package poetry

import (
	"fmt"
	"unicode"
)

type Line string
type Stanza []Line
type Poem []Stanza

func NewPoem() Poem {
	return Poem{}
}

func (p Poem) NumStanzas() int {
	return len(p)
}

func (p Poem) NumLines() (count int) {
	for _, s := range p {
		count += s.NumLines()
	}
	return
}

func (s Stanza) NumLines() int {
	return len(s)
}

func (p Poem) Stats() (numVowels, numConsonants, numPuncs int) {
	for _, s := range p {
		for _, l := range s {
			for _, r := range l {
				if unicode.IsPunct(r) {
					numPuncs++
				} else if unicode.IsSpace(r) {
					// Ignore whitespace
				} else {
					switch r {
					case 'a', 'e', 'i', 'o', 'u':
						numVowels++
					default:
						numConsonants++
					}
				}
			}
		}
	}
	return
}

func (s Stanza) String() (result string) {
	for _, l := range s {
		result += fmt.Sprintf("%s\n", l)
	}
	return
}

func (p Poem) String() (result string) {
	for _, s := range p {
		result += fmt.Sprintf("%s\n", s)
	}
	return
}
