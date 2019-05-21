package esc_test

import (
	"reflect"
	"testing"

	"cmd/compile/internal/esc"
)

func expectEscapes(t *testing.T, loc *esc.Loc, want bool) {
	if got := loc.Escapes(); got != want {
		t.Errorf("%v.Leaks(): got %v, want %v", loc, got, want)
	}
}

func TestBasic(t *testing.T) {
	var s esc.Solver

	a := s.NewLoc(0)
	b := s.NewLoc(0)
	c := s.NewLoc(0)

	s.Assign(&s.Heap, a, -1)
	s.Assign(&s.Heap, b, 0)
	s.Assign(&s.Heap, c, 1)

	s.Solve()

	expectEscapes(t, a, true)
	expectEscapes(t, b, false)
	expectEscapes(t, c, false)
}

func TestRecursive(t *testing.T) {
	var s esc.Solver

	a := s.NewLoc(0)
	b := s.NewLoc(0)

	s.Assign(&s.Heap, a, 0)
	s.Assign(a, b, -1)

	s.Solve()

	expectEscapes(t, a, false)
	expectEscapes(t, b, true)
}

func TestMore(t *testing.T) {
	var s esc.Solver

	a := s.NewLoc(0)
	b := s.NewLoc(0)
	c := s.NewLoc(0)

	s.Assign(&s.Heap, a, 1)
	s.Assign(a, b, -1)
	s.Assign(b, c, -1)

	s.Solve()

	expectEscapes(t, a, false)
	expectEscapes(t, b, false)
	expectEscapes(t, c, true)
}

func TestDepth(t *testing.T) {
	var s esc.Solver

	a := s.NewLoc(0)
	b := s.NewLoc(1)
	c := s.NewLoc(1)

	s.Assign(a, b, 0)
	s.Assign(a, c, -1)

	s.Solve()

	expectEscapes(t, a, false)
	expectEscapes(t, b, false)
	expectEscapes(t, c, true)
}

func TestReturnLeak(t *testing.T) {
	var s esc.Solver

	f := s.NewFunc(0, 1)

	r := f.Results[0]
	a := s.NewLoc(0)
	b := s.NewLoc(0)

	s.Assign(r, a, 0)
	s.Assign(r, b, -1)

	s.Solve()

	expectEscapes(t, r, false)
	expectEscapes(t, a, false)
	expectEscapes(t, b, true)
}

func TestLeakParam(t *testing.T) {
	var x esc.Summary

	{
		var s esc.Solver

		f := s.NewFunc(1, 0)
		p := f.Params[0]

		s.Assign(&s.Heap, p, -1)

		x = s.Solve()

		expectEscapes(t, p, true)
	}

	var s esc.Solver

	a := s.NewLoc(0)
	b := s.NewLoc(0)

	s.Assign(a, b, -1)
	s.Call(x, []*esc.Loc{a}, nil)

	s.Solve()

	expectEscapes(t, a, false)
	expectEscapes(t, b, true)
}

func TestNotes(t *testing.T) {
	var s esc.Solver

	f := s.NewFunc(4, 2)

	p1 := f.Params[0]
	p2 := f.Params[1]
	p3 := f.Params[2]

	r1 := f.Results[0]
	r2 := f.Results[1]

	s.Assign(&s.Heap, p1, 0)
	s.Assign(&s.Heap, p2, 1)
	s.Assign(r1, p2, 2)
	s.Assign(r2, p3, 3)

	x := s.Solve()

	want := []string{"", "esc:0x3a", "esc:0x202", "esc:0x1"}
	got := x.Notes()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Notes: got %v, want %v", got, want)
	}
}

func TestRecursionF8(t *testing.T) {
	var s esc.Solver

	// From test/escape_because.go:
	//
	// func f8(x int, y *int) *int {
	//     if x <= 0 { return y }
	//     x--
	//     return f8(*y, &x)
	// }

	f := s.NewFunc(2, 1)
	p1 := f.Params[0]
	p2 := f.Params[1]
	r1 := f.Results[0]

	s.Assign(r1, p2, 0) // return y

	t0 := s.NewLoc(0)
	t1 := s.NewLoc(0)
	s.Assign(t0, p2, 1)                               // t0 = *y
	s.Assign(t1, p1, -1)                              // t1 = &x
	s.CallFunc(f, []*esc.Loc{t0, t1}, []*esc.Loc{r1}) // return f8(*y, &x)

	s.Solve()

	expectEscapes(t, p1, true)
	expectEscapes(t, p2, false)
	expectEscapes(t, t0, false)
	expectEscapes(t, t1, false)
}
