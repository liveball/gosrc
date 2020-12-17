package main

import (
	"fmt"
	"gosrc/design_patterns/wire/attr/impl"
)

func main()  {
	app:=impl.Init()
    fmt.Println(app.Foo.A)
}
