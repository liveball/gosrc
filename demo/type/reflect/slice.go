package main

import (
	"fmt"
	"reflect"
	"time"
)

func copyInsert(slice interface{}, pos int, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	v = reflect.Append(v, reflect.ValueOf(value))
	reflect.Copy(v.Slice(pos+1, v.Len()), v.Slice(pos, v.Len()))
	v.Index(pos).Set(reflect.ValueOf(value))
	return v.Interface()
}

func Insert(slice interface{}, pos int, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	ne := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(value)), 1, 1)
	ne.Index(0).Set(reflect.ValueOf(value))
	v = reflect.AppendSlice(v.Slice(0, pos), reflect.AppendSlice(ne, v.Slice(pos, v.Len())))
	return v.Interface()
}

func main() {
	t0 := time.Now()
	for i := 1; i < 10000000; i++ {
		slice := []int{1, 2}
		slice = append(slice[:1], append([]int{i}, slice[1:]...)...)
	}
	t1 := time.Now()
	for i := 1; i < 10000000; i++ {
		slice2 := []int{1, 2}
		slice2 = Insert(slice2, 1, i).([]int)
	}

	t2 := time.Now()
	for i := 1; i < 10000000; i++ {
		slice3 := []int{1, 2}
		slice3 = copyInsert(slice3, 1, i).([]int)
		//  fmt.Println(slice3)
	}

	t3 := time.Now()

	fmt.Println("reflect append insert:", t2.Sub(t1), "append insert: ", t1.Sub(t0), "copy Insert: ", t3.Sub(t2))
}
