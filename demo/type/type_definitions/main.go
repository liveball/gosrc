package main

// A builtinId is the id of a builtin function.
type builtinId int

const (
	// universe scope
	_Append builtinId = iota
	_Cap
	_Close
	_Complex
	_Copy
	_Delete
	_Imag
	_Len
	_Make
	_New
	_Panic
	_Print
	_Println
	_Real
	_Recover
	// package unsafe
	_Alignof
	_Offsetof
	_Sizeof
	// testing support
	_Assert
	_Trace
)

// exprKind describes the kind of an expression; the kind
// determines if an expression is valid in 'statement context'.
type exprKind int

const (
	conversion exprKind = iota
	expression
	statement
)

var predeclaredFuncs = [...]struct {
	name     string
	nargs    int
	variadic bool
	kind     exprKind
}{
	_Append:  {"append", 1, true, expression},
	_Cap:     {"cap", 1, false, expression},
	_Close:   {"close", 1, false, statement},
	_Complex: {"complex", 2, false, expression},
	_Copy:    {"copy", 2, false, statement},
	_Delete:  {"delete", 2, false, statement},
	_Imag:    {"imag", 1, false, expression},
	_Len:     {"len", 1, false, expression},
	_Make:    {"make", 1, true, expression},
	_New:     {"new", 1, false, expression},
	_Panic:   {"panic", 1, false, statement},
	_Print:   {"print", 0, true, statement},
	_Println: {"println", 0, true, statement},
	_Real:    {"real", 1, false, expression},
	_Recover: {"recover", 0, false, statement},

	_Alignof:  {"Alignof", 1, false, expression},
	_Offsetof: {"Offsetof", 1, false, expression},
	_Sizeof:   {"Sizeof", 1, false, expression},

	_Assert: {"assert", 1, false, statement},
	_Trace:  {"trace", 0, true, statement},
}

func defPredeclaredFuncs() {
	for i := range predeclaredFuncs {
		id := builtinId(i)
		if id == _Assert || id == _Trace {
			continue // only define these in testing environment
		}
		def(newBuiltin(id))
	}
}

// A Builtin represents a built-in function.
// Builtins don't have a valid type.
type Builtin struct {
	name string
	id   builtinId
}

func newBuiltin(id builtinId) *Builtin {
	return &Builtin{name: predeclaredFuncs[id].name, id: id}
}

func def(b *Builtin) {
	println(b.id, b.name)
}

func main() {
	defPredeclaredFuncs()
}
