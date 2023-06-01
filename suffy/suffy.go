package suffy

import (
	"errors"
	"unicode/utf8"
)

// Suffy is a friendly nickname for suffix automaton, this implementation is based on this blog: https://codeforces.com/blog/entry/20861
// but it will use arrays instead of maps for better memory complexity at the cost of time complexity
type Suffy struct {
	edges  []hMap
	link   []int
	length []int
	last   int
}

// New initializes a new Suffy
func New() *Suffy {
	suffy := &Suffy{make([]hMap, 1), make([]int, 1), make([]int, 1), 0}
	suffy.link[0] = -1

	return suffy
}

func (suffy *Suffy) InsertString(s string) error {
	if !utf8.ValidString(s) {
		return errors.New("invalid UTF-8 encoded string")
	}

	if suffy.edges == nil {
		*suffy = *New()
	}

	for _, char := range s {
		suffy.unsafeInsert(char)
	}

	return nil
}

func (suffy *Suffy) Insert(char rune) error {
	if !utf8.ValidRune(char) {
		return errors.New("invalid UTF-8 encoded rune")
	}

	suffy.unsafeInsert(char)

	return nil
}

func (suffy *Suffy) unsafeInsert(char rune) {
	// create a new state for the full string
	suffy.edges = append(suffy.edges, hMap{})
	suffy.length = append(suffy.length, suffy.length[suffy.last]+1)
	suffy.link = append(suffy.link, 0)
	r := len(suffy.edges) - 1

	p := suffy.last
	for p >= 0 {
		// if this state already has a transition through char we stop
		if _, ok := suffy.edges[p].Get(char); ok {
			break
		}
		suffy.edges[p].Insert(char, r)
		p = suffy.link[p]
	}
	if p != -1 {
		q, _ := suffy.edges[p].Get(char)
		if suffy.length[p]+1 == suffy.length[q] {
			suffy.link[r] = q
		} else {
			suffy.edges = append(suffy.edges, suffy.edges[q].Copy())
			suffy.length = append(suffy.length, suffy.length[p]+1)
			suffy.link = append(suffy.link, suffy.link[q])
			qq := len(suffy.edges) - 1
			suffy.link[r] = qq
			suffy.link[q] = qq

			for p >= 0 {
				// if this state already has a transition through char we stop
				if target, ok := suffy.edges[p].Get(char); !ok || target != qq {
					break
				}
				suffy.edges[p].Insert(char, qq)
				p = suffy.link[p]
			}
		}
	}

	suffy.last = r
	return
}

func (suffy *Suffy) IsSubstring(s string) (bool, error) {
	if !utf8.ValidString(s) {
		return false, errors.New("invalid UTF-8 encoded string")
	}

	pos := 0
	for _, char := range s {
		var ok bool
		pos, ok = suffy.edges[pos].Get(char)
		if !ok {
			return false, nil
		}
	}

	return true, nil
}
