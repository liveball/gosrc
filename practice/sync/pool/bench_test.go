package main

import (
	"gosrc/go/src/fmt"
	"testing"
	"time"
)

func TestBuffer(t *testing.T) {
	fmt.Println(time.Now().AddDate(0, 0, -1).Add(-12 * time.Hour).Format("20060102 15:04:05"))

	size := 42
	b := newBuffer(size)
	if cap(b) < size {
		t.Fatalf("Excpected cap >= %d, got %d", size, cap(b))
	}
	if len(b) != size {
		t.Fatalf("Excpected len %d, got %d", size, len(b))
	}
	freeBuffer(b)

	size = 40
	b = newBuffer(size)
	if cap(b) < size {
		t.Fatalf("Excpected cap >= %d, got %d", size, cap(b))
	}
	if len(b) != size {
		t.Fatalf("Excpected len %d, got %d", size, len(b))
	}
	freeBuffer(b)

	size = 45
	b = newBuffer(size)
	if cap(b) < size {
		t.Fatalf("Excpected cap >= %d, got %d", size, cap(b))
	}
	if len(b) != size {
		t.Fatalf("Excpected len %d, got %d", size, len(b))
	}
	freeBuffer(b)
}