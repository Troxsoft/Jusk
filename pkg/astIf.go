package pkg

import "fmt"

type Boolean struct {
	Bool bool
}

func (B Boolean) Kind() int {
	return TypeBoolean
}
func (B Boolean) Value() any {
	return B.Bool
}

type IFCondition struct {
	OP   *BinaryExpression
	Bool *Boolean
	Body BodyStatement
}

func (i IFCondition) Kind() int {
	return TypeIf
}
func (i IFCondition) Value() any {
	return nil
}
func (Cpp CppCode) Value() any {
	return Cpp.Code
}

// comprobar y debuggear si funciona
func (a *Ast) parseIf() IFCondition {
	a.next()
	e := a.parseStmt(false)
	if e.Kind() != TypeBinaryExpression && e.Kind() != TypeBoolean {
		panic(fmt.Sprintf("Expectative logical operations(==,<,>,>=,<=,true,false) but found: %+v", e))
	}

	if e.Kind() == TypeBoolean {
		p := a.parseStmt(false)
		if p.Kind() != TypeBody {
			panic(fmt.Sprintf("Expectative if body but found: %+v", p))
		}
		k := e.(Boolean)

		return IFCondition{
			OP:   nil,
			Bool: &k,
			Body: p.(BodyStatement),
		}
	} else {
		p := a.parseStmt(false)
		if p.Kind() != TypeBody {
			panic(fmt.Sprintf("Expectative if body but found: %+v", p))
		}
		k := e.(BinaryExpression)
		return IFCondition{
			OP:   &k,
			Bool: nil,
			Body: p.(BodyStatement),
		}

	}

}
