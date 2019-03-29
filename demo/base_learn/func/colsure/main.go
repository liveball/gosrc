package main

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

type Options struct {
	Name string
	Age  int
}

type Option func(*Options)

func main() {
	opts := newOptions(Name("bbb"), Name("ccc"), Age(11))
	spew.Dump(opts)
}

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

func Age(n int) Option {
	return func(o *Options) {
		o.Age = n
	}
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Name: "aaa",
	}

	for k, o := range opts {
		o(&opt)
		spew.Dump(k, opt, reflect.TypeOf(o).Kind())
	}

	return opt
}
