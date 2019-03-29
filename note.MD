
### 研究源码包  src/go/... 

### 代码
``` go
    package main

    import (
        "fmt"
    )

    type Stringer interface {
        String()
        test()
    }

    // type Celsius float64

    // func (c Celsius) String() string {
    // 	return strconv.FormatFloat(float64(c), 'f', 1, 64) + "C"
    // }

    type Day int

    var dayName = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}

    func (day Day) String() string {
        return dayName[day]
    }
    func (day Day) test() string {
        return dayName[4]
    }

    func main() {
        fmt.Println(Day(0))
        // fmt.Println(Day(1), "温度是:", Celsius(20.6))
    }

```
### 问题：fmt.Println(Day(0)) 自动调用 String()

### 结论：
    fmt.Println(Day(0)) 为啥调用了String()，原因 Day(0) 被断言：v := Day(0).(type)为Stringer，
然后执行v.String()并把执行的结果送入buf.WriteString，如果你把它的这个内部处理干掉就不会了，fmt.Println(Day(0)) 只会输出0


### 问题: type xxx underlying(int/float/string/byte/) 啥意思？怎么实现类型解析？如果绑定方法，方法是如何绑定的？

importer.Default()

``` go
// Default returns an Importer for the compiler that built the running binary.
// If available, the result implements types.ImporterFrom.
func Default() types.Importer {
	return For(runtime.Compiler, nil)
}

func For(compiler string, lookup Lookup) types.Importer {
	switch compiler {
	case "gc":
		return &gcimports{
			packages: make(map[string]*types.Package),
			lookup:   lookup,
        }
        ...

pkg, err := conf.Check("temperature", fset, []*ast.File{f}, nil)


// Files checks the provided files as part of the checker's package.
func (check *Checker) Files(files []*ast.File) error { return check.checkFiles(files) }

func (check *Checker) checkFiles(files []*ast.File) (err error) {
	defer check.handleBailout(&err)
	check.initFiles(files)

	check.collectObjects()

	check.packageObjects()// pkg 

	check.processDelayed(0) // incl. all functions

	check.initOrder()

	if !check.conf.DisableUnusedImportCheck {
		check.unusedImports()
	}

	check.recordUntyped()

	check.pkg.complete = true
	return
}

//对象类型 - 检查其各自（文件）上下文中的obj声明。
//有关def和path的详细信息，请参阅check.type。
func (check *Checker) objDecl(obj Object, def *Named, path []*TypeName) {

//  - 类型尚未知的对象被涂成白色（初始颜色）
//  - 类型处于推断过程中的对象被涂成灰色
//  - 完全推断出类型的对象被涂成黑色

// stmt typechecks statement s.
func (check *Checker) stmt(ctxt stmtContext, s ast.Stmt) {}
func (check *Checker) builtin(x *operand, call *ast.CallExpr, id builtinId) (_ bool) {
func (check *Checker) funcBody(decl *declInfo, name string, sig *Signature, body *ast.BlockStmt, iota constant.Value) 

1、types.Config{Importer: importer.Default()}.Check()->Files()->heckFiles()->packageObjects()->objDecl()->funcDecl()->funcBody() 

2、call()->builtin()->rawExpr()-> exprInternal()->funcBody() check.typExpr(e.Type, nil, nil)

stmtList()
stmt(inner, s)
declStmt()
varDecl()
typExpr()
typExprInternal()
selector()

```

``` go

// parse
func ParseFile(fset *token.FileSet, filename string, src interface{}, mode Mode) (f *ast.File, err error) 

func (p *parser) parseFile() *ast.File

func (p *parser) parseDecl(sync map[token.Token]bool) ast.Decl 

func (p *parser) parseFuncDecl() *ast.FuncDecl

func (p *parser) parseParameters(scope *ast.Scope, ellipsisOk bool) *ast.FieldList 

func (p *parser) parseParameterList(scope *ast.Scope, ellipsisOk bool) (params []*ast.Field)

func (p *parser) resolve(x ast.Expr)

func (p *parser) tryResolve(x ast.Expr, collectUnresolved bool)

panic.go
func gopanic(e interface{}) {
     reflectcall(nil, unsafe.Pointer(d.fn), deferArgs(d), uint32(d.siz), uint32(d.siz))
}

parseTypeDefinition()

// parse 
parseUnaryExpr()
parsePrimaryExpr()
parseOperand()
parseFuncTypeOrLit()
parseBody()
parseCaseClause()
parseStmtList()
parseStmt()
parseForStmt()
parseSimpleStmt()
parseLhsList()
parseExprList()
parseExpr()
parseBinaryExpr()

```

### ssa ？
https://github.com/dominikh/go-tools/blob/master/ssa/builder.go