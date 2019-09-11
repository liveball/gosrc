package syscall2

import "unsafe"

func Read(fd int32, p unsafe.Pointer, n int32) int32
