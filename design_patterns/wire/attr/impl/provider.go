package impl

import "github.com/google/wire"

type Foo struct {
	A int
}

type Bar struct {
	B int
	C string
}

type App struct {
	Foo *Foo
	Bar *Bar
}

func DefaultApp(foo *Foo, bar *Bar) *App {
	return &App{Foo: foo, Bar: bar}
}

func NewFoo() *Foo{
	return &Foo{A:100}
}

var provideFoo = wire.NewSet(NewFoo) // -> inject only App.Foo
//var provideBar = wire.NewSet(wire.Struct(new(Bar), "*")) //-> inject all fields
