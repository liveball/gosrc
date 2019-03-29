package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	// Parse a single source file.
	const input = `
package temperature
import "fmt"
type Celsius float64
var A int
func (c Celsius) String() string  { return fmt.Sprintf("%g°C", c) }
func (c *Celsius) SetF(f float64) { *c = Celsius(f - 32 / 9 * 5) }
type foo struct {
}

func main() {
	judge(foo{})
	// judge(&foo{})
}

//assert
func judge(i interface{}) {
	switch i.(type) {
	case struct{}:
		fmt.Println("struct")
	case *struct{}:
		fmt.Println("*struct")
	default:
		fmt.Println("no")
	}
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "celsius.go", input, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Type-check a package consisting of this file.
	// Type information for the imported packages
	// comes from $GOROOT/pkg/$GOOS_$GOOARCH/fmt.a.
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("temperature", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print the method sets of Celsius and *Celsius.
	celsius := pkg.Scope().Lookup("Celsius").Type()
	// fmt.Println(celsius.String(), celsius.Underlying()) //temperature.Celsius float64

	for _, t := range []types.Type{celsius, types.NewPointer(celsius)} {
		// fmt.Printf(">> Method set of %v:\n", t)
		mset := types.NewMethodSet(t)
		for i := 0; i < mset.Len(); i++ {
			// mset.At(i)
			// fmt.Println(mset.At(i))
		}
		fmt.Println()
	}

}
