/*
Package poetry package demostrates the basics of Golang, including types, slices, interfaces, strings, fmt, bufio, etc.
It's from John Graham-Cumming's Introduction to Go Programming course.
*/
package poetry

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A Line is the basic building block of a Poem
type Line string

// A Stanza is a slice of Lines
type Stanza []Line

// A Poem is a slice of Stanzas
type Poem []Stanza

// NewPoem returns an empty Poem
func NewPoem() Poem {
	return Poem{}
}

// LoadPoem loads text from a file, builds a Poem, and returns it
func LoadPoem(name string) (Poem, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := Poem{}

	var s Stanza

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		l := scan.Text()
		if l == "" {
			p = append(p, s)
			s = Stanza{}
			continue
		}

		s = append(s, Line(l))
	}
	p = append(p, s)

	if scan.Err() != nil {
		return nil, scan.Err()
	}

	return p, nil
}

// to use sort.Sort on a Stanza, we need to implement Len(), Swap(), and Less()
func (s Stanza) Len() int {
	return len(s)
}

// to use sort.Sort on a Stanza, we need to implement Len(), Swap(), and Less()
func (s Stanza) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// to use sort.Sort on a Stanza, we need to implement Len(), Swap(), and Less()
func (s Stanza) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// NumLines returns the number of Lines in a Stanza
func (s Stanza) NumLines() int {
	return len(s)
}

// NumStanzas returns the number of Stanzas in a Poem
func (p Poem) NumStanzas() int {
	return len(p)
}

// NumWords returns the number of words in a Poem
func (p Poem) NumWords() int {
	count := 0
	for _, s := range p {
		for _, l := range s {
			sl := string(l)
			parts := strings.Split(sl, " ")
			count += len(parts)
		}
	}
	return count
}

// NumThe returns the number of lines that contains "The" in a Poem
func (p Poem) NumThe() int {
	count := 0
	for _, s := range p {
		for _, l := range s {
			sl := string(l)
			if strings.Contains(sl, "The") {
				count++
			}
		}
	}
	return count
}

// NumLines returns the number of lines in a Poem
// This function demostrates the usages of named returned value
func (p Poem) NumLines() (count int) {
	for _, s := range p {
		count += s.NumLines()
	}
	return
}

// Stats returns three integers for a Poem: the number of vowels, consonants, and punctuations
func (p Poem) Stats() (numVowels, numConsonants, numPuncs int) {
	for _, s := range p {
		for _, l := range s {
			for _, r := range l {
				switch r {
				case 'a', 'e', 'i', 'o', 'u':
					numVowels++
				case ',', ' ', '!':
					numPuncs++
				default:
					numConsonants++
				}
			}
		}
	}
	return
}

func (s Stanza) String() string {
	var result string
	for _, l := range s {
		result += fmt.Sprintf("%s\n", l)
	}
	return result
}

func (p Poem) String() string {
	var result string
	for _, s := range p {
		result += fmt.Sprintf("%s\n", s)
	}
	return result
}
