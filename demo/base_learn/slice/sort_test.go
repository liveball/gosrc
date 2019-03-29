package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type total []int64

func (t total) Len() int           { return len(t) }
func (t total) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t total) Less(i, j int) bool { return t[i] > t[j] }

func TestSort(t *testing.T) {
	m := []int64{789, 123, 456}
	// sort.Sort(total(t))
	sort.Slice(m, func(i, j int) bool {
		return m[i] < m[j]
	})
	fmt.Printf("%v\n", m)
}

func TestMapsort(t *testing.T) {
	ps := make([]*Person, 0)

	p := &Person{
		ID:   1,
		Rank: 2,
	}
	ps = append(ps, p)

	p1 := &Person{
		ID:   2,
		Rank: 1,
	}
	ps = append(ps, p1)

	p2 := &Person{
		ID:   3,
		Rank: 0,
	}
	ps = append(ps, p2)

	p3 := &Person{
		ID:   4,
		Rank: 0,
	}
	ps = append(ps, p3)
	// spew.Dump(ps)

	ps0 := make([]*Person, 0)
	ps1 := make([]*Person, 0)
	ps2 := make([]*Person, 0)

	for _, v := range ps {
		switch v.Rank {
		case 0:
			ps0 = append(ps0, v)
		case 1:
			ps1 = append(ps1, v)
		case 2:
			ps2 = append(ps2, v)
		}

	}

	pps := append(append(ps0, ps1...), ps2...)

	spew.Dump(pps)
}
