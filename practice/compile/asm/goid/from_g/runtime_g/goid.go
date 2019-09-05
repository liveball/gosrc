package runtime_g

import "reflect"

func getg() interface{}

func GetGoid() int64 {
	g := getg()
	gid := reflect.ValueOf(g).FieldByName("goid").Int()
	return gid
}

