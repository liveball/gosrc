package esc

import "fmt"

// Implements compatibility with cmd/compile's legacy esc notes
// system.

// These are both EscContentEscapes:
//
//     func F(x **int) { global = *x }
//     func G(x ***int) { global = **x }
//
// That is, if there's a path from the parameter to sink with 1 or
// more dereferences, then EscContentEscapes is set.

// If the value leaks directly to sink, then we can encode simply as
// EscUnknown.

// cmd/compile encodes paths from params to results even when they're
// longer than the path to sink, unless the path to sink is 0.

func FromNotes(tags []string) Summary {
	panic("TODO")
}

func (s *Summary) Notes() []string {
	notes := make([]string, s.params)
	for i := range notes {
		notes[i] = s.note(i)
	}
	return notes
}

func (s *Summary) note(pi int) string {
	if s.derefs[0][pi] == 0 {
		return ""
	}

	var e uint16

	for ri, derefs := range s.derefs[1:] {
		x := derefs[pi]
		if x >= inf {
			continue
		}

		if x > maxEncodedLevel {
			x = maxEncodedLevel
		}

		e |= uint16(1+x) << uint(3*ri+4)
	}

	if e != 0 {
		e |= escReturn
	} else {
		e |= escNone
	}

	if s.derefs[0][pi] > 0 && s.derefs[0][pi] < inf {
		e |= escContentEscapes
	}

	return fmt.Sprintf("esc:0x%x", e)
}

const (
	escNone           = 0x1
	escReturn         = 0x2
	escContentEscapes = 0x8

	maxEncodedLevel = 6
)
