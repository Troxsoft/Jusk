package pkg

import "fmt"

func (c *Compiler) getVar(name string, scope *ScopeInfoCompile, pro *ProgramInfoCompile) *VarInfoCompile {
	for _, v := range pro.globalScope.vars {
		if v.Name == name {
			return v
		}
	}

	for _, v := range scope.vars {
		if v.Name == name {
			return v
		}
	}
	return nil
}
func (c *Compiler) toCppVarAssing(v AssingDeclaration, pro *ProgramInfoCompile, scope *ScopeInfoCompile, afterPkg bool) string {

	return fmt.Sprintf("%s = %s;", v.Symbol.(Identify).Val.(string), c.GenCode(v.Val.(Stmt), true, pro, scope, afterPkg))
}

func (c *Compiler) toCppVar(v VarDeclaration, pro *ProgramInfoCompile, scope *ScopeInfoCompile, afterPkg bool) string {
	if !v.TypeOnCompiled {
		return fmt.Sprintf("%s %s=%s;", replaceTypesPrimitivesForCppType(v.Type.(AssingType).Type.(string)), v.Symbol.(Identify).Val.(string), c.GenCode(v.Val.(Stmt), false, pro, scope, afterPkg))
	} else {
		return fmt.Sprintf("%s %s=%s;", replaceTypesPrimitivesForCppType(v.Type.(AssingType).Type.(string)), v.Symbol.(Identify).Val.(string), c.GenCode(v.Val.(AssingDeclaration).Val.(Stmt), false, pro, scope, afterPkg))

	}
}
