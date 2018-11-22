package main

import (
	"encoding/json"
	"fmt"
	"go/token"
	"go/types"
	"log"
	"testing"
)

func TestType(t *testing.T) {
	testIsAlias()
}

func testIsAlias() {
	check := func(obj *types.TypeName, want bool) {
		if got := obj.IsAlias(); got != want {
			log.Printf("%v: got IsAlias = %v; want %v", obj, got, want)
		}
	}
	// predeclared types
	check(types.Unsafe.Scope().Lookup("Pointer").(*types.TypeName), false)

	for _, name := range types.Universe.Names() {
		println(name)
		if obj, _ := types.Universe.Lookup(name).(*types.TypeName); obj != nil {
			// fmt.Printf("obj(%#v) \n", reflect.ValueOf(obj))
			check(obj, name == "byte" || name == "rune")

		}
	}

	// various other types
	pkg := types.NewPackage("/data/app/go/src/go1.11.1/demo/type/define", "p")
	t1 := types.NewTypeName(0, pkg, "t1", nil)
	n1 := types.NewNamed(t1, new(types.Struct), nil)
	for _, test := range []struct {
		name  *types.TypeName
		alias bool
	}{
		{types.NewTypeName(0, nil, "t0", nil), false}, // no type yet
		{types.NewTypeName(0, pkg, "t0", nil), false}, // no type yet
		{t1, false}, // type name refers to named type and vice versa
		// {types.NewTypeName(0, nil, "t2", &emptyInterface), true},            // type name refers to unnamed type
		{types.NewTypeName(0, pkg, "t3", n1), true},                         // type name refers to named type with different type name
		{types.NewTypeName(0, nil, "t4", types.Typ[types.Int32]), true},     // type name refers to basic type with different name
		{types.NewTypeName(0, nil, "int32", types.Typ[types.Int32]), false}, // type name refers to basic type with same name
		{types.NewTypeName(0, pkg, "int32", types.Typ[types.Int32]), true},  // type name is declared in user-defined package (outside Universe)
		{types.NewTypeName(0, nil, "rune", types.Typ[types.Rune]), true},    // type name refers to basic type rune which is an alias already
	} {
		// log.Printf("name (%v): alias(%v)", test.name, test.alias)
		check(test.name, test.alias)
	}

	// println(pkg.Scope().Lookup("T").Type().Underlying())
}

var fset = token.NewFileSet()

type H5Recommend struct {
	Tag map[string]map[string][]*TagNameForTypeMap
}

//TagNameForTypeMap tag name map arc type.
type TagNameForTypeMap struct {
	Key []int
	Val []int
}

func TestJson(t *testing.T) {
	var h5 map[string]map[string]interface{}
	err := json.Unmarshal([]byte(recommendMap), &h5)
	if err != nil {
		t.Logf("error(%v)", err)
	}
	fmt.Printf("h5(%+v)", h5)
}

var (
	//RecommendMap 个性化推荐，推荐标题根据对应最近投稿分区确定推荐目录名称
	recommendMap = `
	  {
		"course":{
			"shoot":{
				"key":[1],
				"val":[1,2,3]
			}
		}
	  }
	`
)
