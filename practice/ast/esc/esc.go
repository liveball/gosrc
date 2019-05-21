package esc

// Escape analysis.
//
// Driver is responsible for decomposing source programs into
// locations, and establishing assignments between those locations.
//
// Solver determines which locations escape to the heap, and also
// computes a summary that describes the escape relationship of
// parameters to results, which can be used to handle function calls
// to opaque (but known) functions.

// TODO(mdempsky): Add NewSolver function? Maybe NewSolver(params, results int) (*Solver, []*Loc, []*Loc)?
// Need to handle solving multiple mutually recursive functions simultaneously, and being able to
// identify these cross-calls.

// TODO(mdempsky): Make Heap private?

// TODO(mdempsky): Variadic functions?

const inf = 1e9

type Solver struct {
	Heap    Loc
	params  []*Loc
	results []*Loc
	all     []*Loc
}

type flow struct {
	src    *Loc
	derefs int
}

type Loc struct {
	flows []flow

	param    bool
	result   bool
	escapes  bool
	depth    int
	distance int
}

type Func struct {
	Params  []*Loc
	Results []*Loc
}

func (s *Solver) NewFunc(params, results int) *Func {
	f := new(Func)
	for i := 0; i < params; i++ {
		f.Params = append(f.Params, s.newParam())
	}
	for i := 0; i < results; i++ {
		f.Results = append(f.Results, s.newResult())
	}
	return f
}

func (s *Solver) newParam() *Loc {
	loc := &Loc{
		param: true,
	}
	s.params = append(s.params, loc)
	s.all = append(s.all, loc)
	return loc
}

func (s *Solver) newResult() *Loc {
	loc := &Loc{
		result: true,
	}
	s.results = append(s.results, loc)
	s.all = append(s.all, loc)
	return loc
}

func (s *Solver) NewLoc(depth int) *Loc {
	if depth < 0 {
		panic("negative depth nonsense")
	}
	loc := &Loc{
		depth: depth,
	}
	s.all = append(s.all, loc)
	return loc
}

func (s *Solver) Assign(dst, src *Loc, derefs int) {
	if derefs < -1 {
		panic("not allowed")
	}
	if src == &s.Heap {
		panic("assigning heap to something is nonsense")
	}

	// outer = &inner
	//
	// TODO(mdempsky): Do we really need to handle this? Should
	// this be the driver's responsibility?
	if derefs < 0 && dst.depth < src.depth {
		dst = &s.Heap
	}

	// TODO(mdempsky): Deduplicate redundant assignments. In
	// particular, for any (dst, src) tuple, I think we only need
	// to track the least derefs value.
	dst.flows = append(dst.flows, flow{src, derefs})
}

func (s *Solver) CallFunc(fn *Func, params, results []*Loc) {
	if len(params) != len(fn.Params) || len(results) != len(fn.Results) {
		panic("parameter list mismatch")
	}

	// Simply wire parameters and results together.
	for i, param := range params {
		s.Assign(fn.Params[i], param, 0)
	}
	for i, result := range results {
		s.Assign(result, fn.Results[i], 0)
	}
}

func (s *Solver) Call(fn Summary, params, results []*Loc) {
	// If we don't have a summary, then just assume parameters all
	// leak to heap.
	if fn.derefs == nil {
		for _, param := range params {
			s.Assign(&s.Heap, param, 0)
		}
		return
	}

	for ri, xx := range fn.derefs {
		for pi, derefs := range xx {
			if derefs >= inf {
				continue
			}

			result := &s.Heap
			if ri > 0 {
				result = results[ri]
			}
			s.Assign(result, params[pi], derefs)
		}
	}
}

func (s *Solver) Solve() Summary {
	derefs := make([][]int, 1+len(s.results))
	for i, result := range s.results {
		derefs[1+i] = s.flood0(result)
	}

	derefs[0] = s.flood0(&s.Heap)

	// TODO(mdempsky): If derefs[1+i][j] > derefs[0][j], then
	// might as well ignore derefs[1+i][j]?

	return Summary{params: len(s.params), derefs: derefs}
}

func (s *Solver) flood0(root *Loc) []int {
	for _, loc := range s.all {
		loc.distance = inf
	}

	root.distance = 0
	s.flood(root)

	out := make([]int, len(s.params))
	for i, param := range s.params {
		x := param.distance
		if x < 0 {
			x = 0
		}
		out[i] = x
	}
	return out
}

func (s *Solver) flood(loc *Loc) {
	base := loc.distance
	if base < 0 {
		base = 0
		loc.escapes = true
	}
	recurse := make([]bool, len(loc.flows))
	for i, flow := range loc.flows {
		dist := base + flow.derefs
		if dist < flow.src.distance {
			flow.src.distance = dist
			recurse[i] = true
		}
	}
	for i, b := range recurse {
		if b {
			s.flood(loc.flows[i].src)
		}
	}
}

func (l *Loc) Escapes() bool {
	return l.escapes
}

// A Summary compactly describes the escape analysis relationships
// between a function's parameters, results, and the heap.
type Summary struct {
	params int

	// derefs[result][param] is the number of derefs separating
	// result from param. As a special case, result 0 actually
	// represents the heap.
	derefs [][]int
}
