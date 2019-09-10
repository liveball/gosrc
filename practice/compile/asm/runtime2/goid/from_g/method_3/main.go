package main

import (
	"reflect"
	"runtime"
	"strings"
)


func getg() interface{}

func GetGoid() int64 {
	g := getg()
	gid := reflect.ValueOf(g).FieldByName("goid").Int()
	return goid
}

func main() {
	println(GetGoid())
}
