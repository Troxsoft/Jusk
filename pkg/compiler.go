package pkg

import (
	"fmt"
)

type Compiler struct {
	Astes    *Ast
	Cpp      string
	Programs []*ProgramInfoCompile
	Public   *ProgramInfoCompile
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

func NewCompiler(ast *Ast) *Compiler {
	return &Compiler{
		Astes:    ast,
		Cpp:      "",
		Programs: []*ProgramInfoCompile{},
		Public:   &ProgramInfoCompile{},
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
		code += fmt.Sprintf("#ifndef _%s_\n", h.PkgName)
		code += fmt.Sprintf("#define _%s_\n", h.PkgName)
		code += "\nnamespace " + h.PkgName + "  {\n"
		for _, v := range h.Body {

			code += com.GenCode(v, true, program, program.globalScope)
		}
		code += "\n}\n"
		code += "\n#endif\n"
		com.Cpp = code
	} else {
		code += fmt.Sprintf("#ifndef _%s_\n", h.PkgName)
		code += fmt.Sprintf("#define _%s_\n", h.PkgName)
		for _, v := range h.Body {

			code += com.GenCode(v, true, program, program.globalScope)
		}
		code += "\n#endif\n"
		com.Cpp = code
	}

}

func (com *Compiler) GenCode(h Stmt, r bool, pro *ProgramInfoCompile, scope *ScopeInfoCompile) string {
	if h.Kind() == TypeCPP {
		o := h.(CppCode)
		return o.Code
	} else if h.Kind() == TypeParent {
		return fmt.Sprintf("(%s)", com.GenCode(h.(Parent).Children[0], true, pro, scope))
	} else if h.Kind() == TypeLiteralNumber {
		return fmt.Sprint(h.(LiteralNumeric).Val)
	} else if h.Kind() == TypeLiteralString {
		return fmt.Sprintf(" \"%s\" ", h.(LiteralString).Val)
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
