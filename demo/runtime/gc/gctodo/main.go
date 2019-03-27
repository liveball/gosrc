package main

import (
	"fmt"
	"go/types"
	"sort"

	"golang.org/x/tools/go/loader"
)

var arches = []string{
	"amd64",
	"arm",
	"arm64",
	"mips",
	"mips64",
	"ppc64",
	"s390x",
	"x86",
}

func main() {
	var cfg loader.Config
	for _, arch := range arches {
		cfg.Import("cmd/compile/internal/" + arch)
	}

	prog, err := cfg.Load()
	if err != nil {
		panic(err)
	}

	m := make(map[types.Object]int)
	note := func(obj types.Object) {
		if pkg := obj.Pkg(); pkg == nil || pkg.Path() != "cmd/compile/internal/gc" {
			return
		}

		switch obj := obj.(type) {
		case *types.Func:
			if obj.Type().(*types.Signature).Recv() != nil {
				return
			}
		case *types.Var:
			if obj.IsField() {
				return
			}
		}

		// spew.Dump(obj)
		m[obj] += 1
	}

	for _, pkg := range prog.InitialPackages() {
		for _, tv := range pkg.Types {
			if named, ok := tv.Type.(*types.Named); ok {
				note(named.Obj())
			}
		}
		for _, obj := range pkg.Uses {
			note(obj)
		}
	}

	var out []string
	for obj, count := range m {
		out = append(out, fmt.Sprintln(obj.Name(), count))
	}
	sort.Strings(out)
	for _, s := range out {
		fmt.Print(s)
	}
}
