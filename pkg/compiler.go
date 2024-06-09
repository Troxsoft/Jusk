package pkg

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type ImportCompiler struct {
	Name         string
	OtherImports []*ImportCompiler
	Compilers    *Compiler
}
type Compiler struct {
	Astes    *Ast
	Cpp      string
	Programs []*ProgramInfoCompile
	Public   *ProgramInfoCompile
	Path     string
	Imports  []*ImportCompiler
}
type ProgramInfoCompile struct {
	scopes      []*ScopeInfoCompile
	globalScope *ScopeInfoCompile
	funcs       []*FunctionInfoCompile
	structs     []*StructInfoCompile
	//NameProgram string
}
type VarInfoCompile struct {
	Type string
	Name string
}
type VarStructInfoCompile struct {
	Type   string
	Name   string
	Public bool
}
type ScopeInfoCompile struct {
	vars []*VarInfoCompile
}
type FunctionInfoCompile struct {
	ReturnType string
	Name       string
	// params types
	Params []string
	Public bool
}
type StructInfoCompile struct {
	Name  string
	Funcs []FunctionInfoCompile
	Vars  []VarStructInfoCompile
	Cpps  []CppCode
}

func NewCompiler(ast *Ast, path string) *Compiler {
	return &Compiler{
		Astes:    ast,
		Cpp:      "",
		Programs: []*ProgramInfoCompile{},
		Public:   &ProgramInfoCompile{},
		Path:     path,
		Imports:  []*ImportCompiler{},
	}
}
func (com *Compiler) Compile() {
	h := com.Astes.Nodes.(ProgramBody)
	code := ""
	program := &ProgramInfoCompile{
		scopes: []*ScopeInfoCompile{},
		globalScope: &ScopeInfoCompile{
			vars: []*VarInfoCompile{},
		},
		funcs:   []*FunctionInfoCompile{},
		structs: []*StructInfoCompile{},

		//NameProgram: v.Kind().,
	}

	if h.PkgName != "main" {
		code += fmt.Sprintf("#ifndef _%s_JK\n", h.PkgName)
		code += fmt.Sprintf("#define _%s_JK\n", h.PkgName)
		code += "#include <iostream>\n#include <string>\n"

		code += "\nnamespace " + h.PkgName + "  {\n"
		for _, v := range h.Body {

			code += com.GenCode(v, true, program, program.globalScope)
		}
		code += "\n}\n"
		code += "\n#endif\n"
		com.Cpp = code
	} else {
		code += "#include <iostream>\n#include <string>\n"
		for _, v := range h.Body {

			code += com.GenCode(v, true, program, program.globalScope)
		}

		com.Cpp = code
	}

}

func (com *Compiler) GenCode(h Stmt, r bool, pro *ProgramInfoCompile, scope *ScopeInfoCompile) string {

	if h.Kind() == TypeCPP {
		o := h.(CppCode)
		o.Code = strings.ReplaceAll(o.Code, "\n", `\n`)
		o.Code = strings.ReplaceAll(o.Code, "@new_line", "\n")
		return o.Code
	} else if h.Kind() == TypeImport {
		o := h.(Import)
		var dataB []byte
		var data string
		var err error
		var juskNww *Jusk
		var ee string

		if !strings.HasPrefix(o.Val.(string), "@") {

			ee = com.Path + o.Val.(string)
			dataB, err = os.ReadFile(com.Path + o.Val.(string))
			if err != nil {
				panic(err.Error())
			}
			data = string(dataB)
		} else {
			ee, _ = os.Executable()
			ee = strings.ReplaceAll(ee, `\`, "/")
			var pp string

			if runtime.GOOS == "windows" {
				pp = ee[:len(ee)-8]
			} else if runtime.GOOS == "linux" {
				pp = ee[:len(ee)-4]
			}
			//p90000 := strings.Split(o.Val.(string)[1:], "/")

			ee = pp + `lib/` + o.Val.(string)[1:]
			dataB, err = os.ReadFile(ee)
			lip := strings.Split(ee, "/")
			ee = ee[:len(ee)-len(lip[len(lip)-1])]

			if err != nil {
				panic(err.Error())
			}
			data = string(dataB)

		}
		juskNww = NewJuskLang(data)
		err = juskNww.Tokenize()
		if err != nil {
			panic(err.Error())
		}
		err = juskNww.GenerateAst()
		if err != nil {
			panic(err.Error())
		}

		p := juskNww.Compile(ee)
		// for _, v := range juskNww.Compiles.Public.funcs {
		// 	com.Public.funcs = append(com.Public.funcs, v)
		// }
		// for _, v := range juskNww.Compiles.Public.structs {
		// 	com.Public.structs = append(com.Public.structs, v)
		// }

		for _, v := range com.Imports {
			if v.Name == juskNww.Astes.Nodes.(ProgramBody).PkgName {
				panic(fmt.Sprintf("Already import package: %s", juskNww.Astes.Nodes.(ProgramBody).PkgName))

			}
		}
		com.Imports = append(com.Imports, &ImportCompiler{
			Name:         juskNww.Astes.Nodes.(ProgramBody).PkgName,
			OtherImports: juskNww.Compiles.Imports,
			Compilers:    juskNww.Compiles,
		})
		return "\n" + p + "\n"

	} else if h.Kind() == TypePointStmt {
		o := h.(PointStmt)
		var pkg *Compiler = nil
		for _, v := range com.Imports {
			if o.Father.(string) == v.Name {
				pkg = v.Compilers
			}
		}
		if pkg == nil {
			panic(fmt.Sprintf("Unknown SYMBOL %s", o.Father.(string)))
		}
		return fmt.Sprintf("%s::%s", o.Father, pkg.GenCode(o.Children.(Stmt), true, pro, scope))
	} else if h.Kind() == TypeParent {
		return fmt.Sprintf("(%s)", com.GenCode(h.(Parent).Children[0], true, pro, scope))
	} else if h.Kind() == TypeLiteralNumber {
		return fmt.Sprint(h.(LiteralNumeric).Val)
	} else if h.Kind() == TypeIf {
		return com.parseIf(h.(IFCondition), pro, scope)
	} else if h.Kind() == TypeLiteralString {
		donaldTrump := fmt.Sprintf(" \"%s\" ", h.(LiteralString).Val)
		donaldTrump = strings.ReplaceAll(donaldTrump, "\n", "\\n")
		return donaldTrump
	} else if h.Kind() == TypeBinaryExpression {
		o := h.(BinaryExpression)
		l := ""
		l += com.GenCode(o.Left.(Stmt), false, pro, scope)
		if o.Operator == PLUS {
			l += "+"
		} else if o.Operator == MINUS {
			l += "-"
		} else if o.Operator == MULTIPLY {
			l += "*"
		} else if o.Operator == DIVIDE {
			l += "/"
		} else if o.Operator == PORCENT {
			l += "%"
		} else if o.Operator == COMPARE {
			l += "=="
		} else if o.Operator == LESS {
			l += "<"
		} else if o.Operator == GREATER {
			l += ">"
		} else if o.Operator == COMPARE_GREATER {
			l += ">="
		} else if o.Operator == COMPARE_LESS {
			l += "<="
		} else if o.Operator == NOCOMPARE {
			l += "!="
		}
		l += com.GenCode(o.Right.(Stmt), false, pro, scope)
		return l
	} else if h.Kind() == TypeAssingDeclaration {

		o := h.(AssingDeclaration)
		if com.getVar(o.Symbol.(Identify).Val.(string), scope) == nil {
			panic(fmt.Sprintf("%s not exists", o.Symbol.(Identify).Val.(string)))
		}
		return com.toCppVarAssing(o, pro, scope)
	} else if h.Kind() == TypeVarDeclaration {

		o := h.(VarDeclaration)
		if com.getVar(o.Symbol.(Identify).Val.(string), scope) != nil {
			panic(fmt.Sprintf("%s already exists", o.Symbol.(Identify).Val.(string)))
		} else {
			scope.vars = append(scope.vars, &VarInfoCompile{
				Name: o.Symbol.(Identify).Val.(string),
				Type: o.Type.(AssingType).Type.(string),
			})
		}
		return com.toCppVar(o, pro, scope)
	} else if h.Kind() == TypeBody {
		_, p := com.toCppBody(h.(BodyStatement), pro, scope)
		return p
	} else if h.Kind() == TypeStruct {
		o := h.(Struct)
		if com.getStruct(o.Symbol.Val.(string), pro) != nil {
			panic(fmt.Sprintf("%s struct already exists", o.Symbol.Val.(string)))
		}
		infoCompile := &StructInfoCompile{
			Name: o.Symbol.Val.(string),
			Cpps: o.Cpps,
		}
		for _, v := range o.Vars {
			infoCompile.Vars = append(infoCompile.Vars, VarStructInfoCompile{
				Type:   v.Type.Type.(string),
				Name:   v.Symbol.Val.(string),
				Public: v.Public,
			})
		}

		pro.structs = append(pro.structs, infoCompile)
		return com.toCppStruct(*infoCompile, pro, scope)
	} else if h.Kind() == TypeReturn {
		o := h.(Return)
		return com.toCppReturn(o, pro, scope)
	} else if h.Kind() == TypeFunctionCall {
		o := h.(FunctionCall)
		return com.toCppFunctionCall(o, r, pro, scope)
	} else if h.Kind() == TypeBoolean {

		o := h.(Boolean)

		if o.Bool {
			return fmt.Sprintf(" true ")
		} else {
			return fmt.Sprintf(" false ")

		}
	} else if h.Kind() == TypeIdentify {
		if com.getVar(h.(Identify).Val.(string), scope) == nil {
			panic(fmt.Sprintf("%s not exists", h.(Identify).Val.(string)))
		}
		return h.(Identify).Val.(string)
	} else if h.Kind() == TypeFunction {
		o := h.(Function)
		fun := &FunctionInfoCompile{
			ReturnType: o.Return.Val.(string),
			Name:       o.Symbol.Val.(string),
			Public:     o.Public,
		}
		if fun.Public {
			com.Public.funcs = append(com.Public.funcs, fun)
		} else {
			pro.funcs = append(pro.funcs, fun)
		}
		for _, p := range o.Arguments {
			fun.Params = append(fun.Params, p.Type.Type.(string))
		}
		return com.toCppFunction(o, *fun, pro, scope)
	} else {
		panic("Invalid sentance: " + fmt.Sprint(h.Kind()))
	}

}
