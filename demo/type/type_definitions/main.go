package main

// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct { /* Mutex fields */
}

func (m *Mutex) Lock()   { /* Lock implementation */ }
func (m *Mutex) Unlock() { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of PtrMutex's underlying type *Mutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex.
type PrintableMutex struct {
	Mutex
}

func main() {
	// 	type NewMutex Mutex

	// real()
	// trace()
	assert()
	// recover()
}
