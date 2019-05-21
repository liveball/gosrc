package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	fmt.Println(errMap()) //wrong
	// concurrentErrMap() //right
}

func errMap() (errMap map[int32]error) {
	errMap = make(map[int32]error)
	errMap01 := make(map[int32]error)
	errMap02 := make(map[int32]error)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		err := errors.New("data error 1")
		if err != nil {
			errMap01[int32(1)] = err
			errMap01[int32(2)] = err
			errMap01[int32(3)] = err
		}
		wg.Done()
	}()
	go func() {
		err := errors.New("data error 2")
		if err != nil {
			errMap02[int32(4)] = err
			errMap02[int32(5)] = err
			errMap02[int32(6)] = err
		}
		wg.Done()
	}()
	wg.Wait()

	cpMapFunc := func(mps ...map[int32]error) {
		for _, mp := range mps {
			if len(mp) == 0 {
				continue
			}
			for k, v := range mp {
				errMap[k] = v
			}
		}
	}
	cpMapFunc(errMap01, errMap02)
	return
}

func concurrentErrMap() (errMap map[int32]error) {
	errMap = make(map[int32]error)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		err := errors.New("data error 1")
		if err != nil {
			errMap[int32(1)] = err
			errMap[int32(2)] = err
			errMap[int32(3)] = err
		}
		wg.Done()
	}()
	go func() {
		err := errors.New("data error 2")
		if err != nil {
			errMap[int32(4)] = err
			errMap[int32(5)] = err
			errMap[int32(6)] = err
		}
		wg.Done()
	}()
	wg.Wait()

	return
}
