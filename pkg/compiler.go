package pkg

import "fmt"

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
	//NameProgram string
}
type VarInfoCompile struct {
	Type string
	Name string
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
		funcs: []*FunctionInfoCompile{},

		//NameProgram: v.Kind().,
	}
	for _, v := range h.Body {

		code += com.GenCode(v, program, program.globalScope)
	}
	com.Cpp = code
}

func (com *Compiler) GenCode(h Stmt, pro *ProgramInfoCompile, scope *ScopeInfoCompile) string {
	if h.Kind() == TypeCPP {
		o := h.(CppCode)
		return o.Code
	} else if h.Kind() == TypeLiteralNumber {
		return fmt.Sprint(h.(LiteralNumeric).Val)
	} else if h.Kind() == TypeLiteralString {
		return fmt.Sprintf(" Str(\"%s\") ", h.(LiteralString).Val)
	} else if h.Kind() == TypeBinaryExpression {
		o := h.(BinaryExpression)
		l := ""
		l += com.GenCode(o.Left, pro, scope)
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
		l += com.GenCode(o.Right, pro, scope)
		return l
	} else if h.Kind() == TypeBody {
		_, p := com.toCppBody(h.(BodyStatement), pro, scope)
		return p
	} else if h.Kind() == TypeFunctionCall {
		o := h.(FunctionCall)
		return com.toCppFunctionCall(o, pro, scope)
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
