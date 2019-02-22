package poetry

import (
	"testing"
)

func TestNumStanzas(t *testing.T) {
	p := Poem{}
	if p.NumStanzas() != 0 {
		t.Fatalf("Empty poem is not empty")
	}
	p = Poem{{
		"This is one",
		"This is two",
	}}
	if p.NumStanzas() != 1 {
		t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
	}
}

func TestNumLines(t *testing.T) {
	p := Poem{}
	if p.NumLines() != 0 {
		t.Fatalf("Empty poem is not empty")
	}
	p = Poem{{
		"This is one",
		"This is two",
	}}
	if p.NumLines() != 2 {
		t.Fatalf("Unexpected line count %d", p.NumLines())
	}
}

func TestStats(t *testing.T) {
	p := Poem{}
	v, c, puncs := p.Stats()
	if v != 0 || c != 0 || puncs != 0 {
		t.Fatalf("Bad number of vowels or consonants or punctuation marks (%d %d %d)", v, c, puncs)
	}

	p = Poem{{"Hello"}}
	v, c, puncs = p.Stats()
	if v != 2 || c != 3 || puncs != 0 {
		t.Fatalf("Bad number of vowels or consonants or punctuation marks (%d %d %d)", v, c, puncs)
	}

	p = Poem{{"Hello, World!"}}
	v, c, puncs = p.Stats()
	if v != 3 || c != 7 || puncs != 3 {
		t.Fatalf("Bad number of vowels or consonants or punctuation marks (%d %d %d)", v, c, puncs)
	}
}

func TestNumWords(t *testing.T) {
	p := Poem{}
	if p.NumWords() != 0 {
		t.Fatal("Wrong number of words")
	}

	p = Poem{{"Hello, world!"}}
	if p.NumWords() != 2 {
		t.Fatal("Wrong number of words")
	}
}
