package pkg

import "fmt"

type CppCode struct {
	Code string
}

func (cpp CppCode) Kind() int {
	return TypeCPP
}
func (a *Ast) parseCpp() CppCode {
	a.next()
	f := a.parseStmt()
	if f.Kind() != TypeParent {
		panic(fmt.Sprintf("Expectative body of c++ code but found: %+v", f))
	}

	f30 := f.(Parent)

	return CppCode{
		Code: f30.Children[0].(LiteralString).Val,
	}
}
