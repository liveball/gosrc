package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/constant"
	"go/format"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/tools/go/loader"
)

var fset = token.NewFileSet()

func main() {
	flag.Parse()

	ssaPkg, err := build.Import("cmd/compile/internal/ssa", "", build.FindOnly)
	if err != nil {
		log.Fatal(err)
	}

	srcs, err := filepath.Glob(filepath.Join(ssaPkg.Dir, "gen", "*.go"))
	if err != nil {
		log.Fatal(err)
	}

	var conf loader.Config
	conf.Fset = fset
	conf.CreateFromFilenames("ssagen", srcs...)
	prog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range prog.InitialPackages() {
		for _, file := range pkg.Files {
			fn := fset.File(file.Pos()).Name()
			arch := strings.TrimSuffix(filepath.Base(fn), "Ops.go")

			var edits []edit
			ast.Inspect(file, func(node ast.Node) bool {
				if lit, ok := node.(*ast.CompositeLit); ok {
					symop(pkg, arch, lit, &edits)
				}
				return true
			})
			if len(edits) == 0 {
				continue
			}
			buf, err := ioutil.ReadFile(fn)
			if err != nil {
				log.Fatal(err)
			}
			s := string(buf)
			for j := len(edits) - 1; j >= 0; j-- {
				e := &edits[j]
				s = s[:e.pos] + e.text + s[e.pos:]
			}
			buf, err = format.Source([]byte(s))
			if err != nil {
				log.Fatal(err)
			}
			err = ioutil.WriteFile(fn, buf, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

type edit struct {
	pos  int
	text string
}

type pattern struct {
	rx     string
	effect effect
}

type effect string

const (
	read  effect = "Read"
	write effect = "Write"
	rw    effect = "RdWr"
	addr  effect = "Addr"
	none  effect = "None"
)

var patterns = []pattern{
	{"^(386|AMD64):MOV([BWLQO]|S[SD])load(idx[1248])?$", read},
	{"^(386|AMD64):MOV([BWLQO]|S[SD])store(idx[1248])?$", write},
	{"^(386|AMD64):MOV[BWL][LQ]SXload$", read},
	{"^(386|AMD64):MOV[BWLQ]storeconst(idx[1248])?$", write},
	{"^(386|AMD64):LEA[LQ][1248]?$", addr},

	{"^AMD64:(ADD|SUB|MUL|AND|OR|XOR)([BWLQ]|S[SD])mem$", read},
	{"^AMD64:(XADD|CMPXCHG|AND|OR)[BWLQ]lock$", rw},
	{"^AMD64:XCHG[LQ]$", rw},
	{"^AMD64:MOV[LQ]atomicload$", read},

	{"^(ARM|MIPS|PPC)(64)?:MOV[BHWDFV][UZ]?load$", read},
	{"^(ARM|MIPS|PPC)(64)?:MOV[BHWDFV]addr$", addr},
	{"^(ARM|MIPS|PPC)(64)?:MOV[BHWDFV]store(zero)?$", write},
	{"^(ARM|MIPS|PPC)(64)?:FMOV[SD]load$", read},
	{"^(ARM|MIPS|PPC)(64)?:FMOV[SD]store$", write},
	{"^ARM:CALLudiv$", none},

	{"^PPC64:ADDconst$", addr},

	{"^S390X:(ADD|SUB|MULL|AND|OR|XOR)[DW]?load$", read},
	{"^S390X:MOV[BHWD](BR|Z)?(atomic)?load(idx)?$", read},
	{"^S390X:MOV[BHWD](BR)?(atomic)?store(idx|const)?$", write},
	{"^S390X:FMOV[SD]load(idx)?$", read},
	{"^S390X:FMOV[SD]store(idx)?$", write},
	{"^S390X:MOVDaddr(idx)?$", read},
	{"^S390X:CLEAR$", write},
	{"^S390X:LAAG?$", rw},
	{"^S390X:MVC$", none}, // doesn't actually use Aux
	{"^S390X:STMG?[234]$", write},
	{"^S390X:LoweredAtomic(Cas|Exchange)(32|64)$", rw},

	{"^generic:Arg$", none},
	{"^generic:Addr$", addr},
	{"^generic:Func$", none},
	{"^generic:(Move|Zero)WB$", none}, // sym is the type symbol used for typedmemmove
	{"^generic:FwdRef$", none},
	{"^generic:Var(Def|Kill|Live)$", none}, // specially handled by plive anyway

	// Apologies to von Neumann, but Go acts like a
	// Harvard architecture.
	{"^generic:StaticCall$", none},
	{"^.*:CALLstatic$", none},
}

func symop(pkg *loader.PackageInfo, arch string, lit *ast.CompositeLit, edits *[]edit) {
	tv, ok := pkg.Types[lit]
	if !ok {
		log.Printf("hmm, weird")
		return
	}
	if tv.Type.String() != "ssagen.opData" {
		return
	}

	var name, aux string
	for _, elt := range lit.Elts {
		elt := elt.(*ast.KeyValueExpr)
		switch elt.Key.(*ast.Ident).Name {
		case "name":
			name = constant.StringVal(pkg.Types[elt.Value].Value)
		case "aux":
			aux = constant.StringVal(pkg.Types[elt.Value].Value)
		}
	}
	if !strings.HasPrefix(aux, "Sym") {
		return
	}

	for _, pattern := range patterns {
		if m, err := regexp.MatchString(pattern.rx, arch+":"+name); err != nil {
			log.Fatal(err)
		} else if m {
			pos := lit.Elts[len(lit.Elts)-1].End()

			sep := ", "
			if pos != lit.Rbrace {
				sep = ",\n"
			}

			text := sep + fmt.Sprintf("symEffect: %q", pattern.effect)
			*edits = append(*edits, edit{fset.Position(pos).Offset, text})
			return
		}
	}

	log.Println("found", name, "with aux", aux, "in", arch)
}
