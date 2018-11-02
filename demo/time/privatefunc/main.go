package main

import (
	"C"
	"fmt"

	_ "unsafe"

	"time"

	forceexport "github.com/alangpierce/go-forceexport"
)

//go:linkname time_now time.now
func time_now() (sec int64, nsec int32)

func main() {
	sec, nsec := time_now() //调用内部包
	fmt.Println(sec, nsec)

	time.Sleep(time.Second)

	var timeNow func() (int64, int32)
	err := forceexport.GetFunc(&timeNow, "time.now")
	if err != nil {
		// Handle errors if you care about name possibly being invalid.
		fmt.Println(err)
		return
	}
	// Calls the actual time.now function.
	sec2, nsec2 := timeNow()
	fmt.Println(sec2, nsec2)

	// var syncThrow func() string
	// err = forceexport.GetFunc(&syncThrow, "runtime.sync_throw")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
