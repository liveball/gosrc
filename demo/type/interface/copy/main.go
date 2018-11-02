package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	a := 4

	spew.Dump(a, copy(a))
}

func copy(a interface{}) (b interface{}) {
	typ := reflect.TypeOf(a)
	b = reflect.New(typ).Interface()

	spew.Dump(a, b)

	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(a); err != nil {
		log.Fatalln(err)
	}

	if err := gob.NewDecoder(buf).Decode(b); err != nil {
		log.Fatalln(err)
	}

	buf.Reset()

	fmt.Println(a.(int))
	fmt.Println(*b.(*int))

	return
}
