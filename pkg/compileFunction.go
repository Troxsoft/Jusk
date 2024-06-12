package pkg

import "fmt"

func (c *Compiler) toCppReturn(f90 Return, pro *ProgramInfoCompile, scope *ScopeInfoCompile, afterkg bool) string {
	return fmt.Sprintf("return %s;", c.GenCode(f90.Val.(Stmt), true, pro, scope, afterkg))
}
func (c *Compiler) toCppFunction(f90 Function, fun FunctionInfoCompile, pro *ProgramInfoCompile, scope *ScopeInfoCompile, afterPkg bool) string {
	existsFunction := 0
	for _, v := range pro.funcs {
		if f90.Symbol.Val.(string) == v.Name {
			existsFunction++
		}
	}
	for _, v := range c.Public.funcs {
		if f90.Symbol.Val.(string) == v.Name {
			existsFunction++
		}
	}
	if existsFunction != 1 {
		panic(fmt.Sprintf("Already exists function: %s", f90.Symbol.Val.(string)))
	}
	if f90.Symbol.Val.(string) == "main" && f90.Return.Val.(string) != "Int" {
		panic("Invalid return type on: main function,Int return type is requiered")
	}
	if f90.Symbol.Val.(string) == "main" && f90.Public != true {
		panic("The main function require public visibility")
	}
	scoNew := scope
	for _, v := range f90.Arguments {
		scoNew.vars = append(scoNew.vars, &VarInfoCompile{
			Type: v.Type.Type.(string),
			Name: v.Symbol.Val.(string),
		})
	}
	_, p := c.toCppBody(f90.Body, pro, scoNew, afterPkg)

	h := fmt.Sprintf("%s %s %s %s", replaceTypesPrimitivesForCppType(fun.ReturnType), fun.Name, c.toCppArgs(f90.Arguments), p)

	return h
}
