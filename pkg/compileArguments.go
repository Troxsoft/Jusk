package pkg

func (c *Compiler) toCppArgs(h []Argument) string {
	l := "("
	for i, v := range h {
		if i != len(h)-1 {
			l += replaceTypesPrimitivesForCppType(v.Type.Type.(string)) + " " + v.Symbol.Val.(string) + ","
		} else {
			l += replaceTypesPrimitivesForCppType(v.Type.Type.(string)) + " " + v.Symbol.Val.(string)
		}
	}
	return l + ")"
}
func (c *Compiler) toCppArgsFuncCall(h []Stmt, pro *ProgramInfoCompile, scope *ScopeInfoCompile) string {
	l := "("
	for i, v := range h {
		if i != len(h)-1 {

			l += c.GenCode(v, pro, scope) + " , "
		} else {
			l += c.GenCode(v, pro, scope) + "  "
		}
	}
	return l + ")"
}
