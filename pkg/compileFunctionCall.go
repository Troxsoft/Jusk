package pkg

import "fmt"

func (c *Compiler) toCppFunctionCall(f FunctionCall, r bool, pro *ProgramInfoCompile, scope *ScopeInfoCompile) string {
	if f.Symbol.Val.(string) != "toStr" {
		existsFunction := false
		var fn *FunctionInfoCompile
		for _, v := range pro.funcs {
			if f.Symbol.Val.(string) == v.Name {
				existsFunction = true
				fn = v
			}
		}
		for _, v := range c.Public.funcs {
			if f.Symbol.Val.(string) == v.Name {
				existsFunction = true
				fn = v
			}
		}

		if !existsFunction {
			panic(fmt.Sprintf("Undefined function: %s", f.Symbol.Val.(string)))
		}
		if len(fn.Params) != len(f.Arguments) {
			panic(fmt.Sprintf("invalid arguments on: %s call", fn.Name))
		}

		l := fmt.Sprintf("%s %s", f.Symbol.Val.(string), c.toCppArgsFuncCall(f.Arguments, pro, scope))
		if r {
			l += ";"
		}
		return l
	} else {
		l := fmt.Sprintf("std::to_string%s", c.toCppArgsFuncCall(f.Arguments, pro, scope))
		if r {
			l += ";"
		}

		return l
	}

}
