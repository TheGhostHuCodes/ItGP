package poetry

import (
	"testing"
)

func TestEmptyPoem(t *testing.T) {
	p := Poem{}

	t.Run("NumStanzasReturnsZero", func(t *testing.T) {
		if p.NumStanzas() != 0 {
			t.Fatalf("Empty poem is not empty!")
		}
	})

	t.Run("NumLinesReturnsZero", func(t *testing.T) {
		if p.NumLines() != 0 {
			t.Fatalf("Empty poem is not empty!")
		}
	})

	t.Run("StatsReturnsZero", func(t *testing.T) {
		v, c, puncs := p.Stats()
		if v != 0 || c != 0 || puncs != 0 {
			t.Fatalf("Bad number of vowels,  consonants, or punctuation marks")
		}
	})
}

func TestNonEmptyPoem(t *testing.T) {
	p := Poem{{"And from my pillow, looking forth by light",
		"Of moon or favouring stars, I could behold",
		"The antechapel where the statue stood",
		"of Newton with his prism and silent face,",
		"The marble index of a mind for ever",
		"Voyaging through strange seas of Thought, alone."}}

	t.Run("NumStanzas", func(t *testing.T) {
		if p.NumStanzas() != 1 {
			t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
		}
	})

	t.Run("NumLines", func(t *testing.T) {
		if p.NumLines() != 6 {
			t.Fatalf("Unexpected line count %d", p.NumLines())
		}
	})
}

func TestStats(t *testing.T) {
	t.Run("WithPoemContainingOnlyCharacters", func(t *testing.T) {
		p := Poem{{"Hello"}}
		v, c, puncs := p.Stats()
		if v != 2 || c != 3 || puncs != 0 {
			t.Fatalf("Bad number of vowels,  consonants, or punctuation marks")
		}
	})

	t.Run("WithPoemContainingCharactersWhitespaceAndPunctuation", func(t *testing.T) {
		p := Poem{{"Hello, World!"}}
		v, c, puncs := p.Stats()
		if v != 3 || c != 7 || puncs != 2 {
			t.Fatalf("Bad number of vowels,  consonants, or punctuation marks (%d, %d, %d)",
				v, c, puncs)
		}
	})
}
