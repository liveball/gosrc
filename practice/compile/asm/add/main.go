package main

//go:noinline
func add(a, b int32) (int32, bool) {
	return a + b, true
}

func main() {
	add(10, 32)
}

// echo 'obase=2;137438953482' | bc
//GOOS=linux GOARCH=amd64 go tool compile -S -N -l

// |      +-------------------------+ <-- 32(SP)
// |      |                         |
// G |    |                         |
// R |    |                         |
// O |    | main.main's saved       |
// W |    |     frame-pointer (BP)  |
// S |    |-------------------------| <-- 24(SP)
// |      |      [alignment]        |
// D |    | "".~r3 (bool) = 1/true  | <-- 21(SP)
// O |    |-------------------------| <-- 20(SP)
// W |    |                         |
// N |    | "".~r2 (int32) = 42     |
// W |    |-------------------------| <-- 16(SP)
// A |    |                         |
// R |    | "".b (int32) = 32       |
// D |    |-------------------------| <-- 12(SP)
// S |    |                         |
//   |    | "".a (int32) = 10       |
//   |    |-------------------------| <-- 8(SP)
//   |    |                         |
//   |    |                         |
//   |    |                         |
// \ | /  | return address to       |
//  \|/   |     main.main + 0x30    |
// -      +-------------------------+ <-- 0(SP) (TOP OF STACK)
