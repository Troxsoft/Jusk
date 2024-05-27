package pkg

import "fmt"

type VariableClass struct {
	Symbol Identify
	Type   AssingType
	Public bool
}

func (v VariableClass) Kind() int {
	return TypeVarDeclarationClass
}

type BodyStatement struct {
	Body []Stmt
}

type Class struct {
	Symbol  Identify
	Vars    []VariableClass
	Methods []Function
	Cpps    []CppCode
}

func (c Class) Kind() int {
	return TypeClass
}

// Kind implements Stmt.
func (b BodyStatement) Kind() int {
	return TypeBody
}
func (a *Ast) parseVariableClass() VariableClass {
	p := a.actual()
	public := true
	if p.Type == PUBLIC {
		public = true
	} else {
		public = false
	}
	a.next()
	name := a.actual()

	if name.Type != SYMBOL {
		panic(fmt.Sprintf("Expectative variable name but found: %+v", name))
	}
	a.next()
	k := a.actual()
	if k.Type != TWO_POINTS {
		panic(fmt.Sprintf("Expectative \":\" but found: %+v", k))
	}
	a.next()
	typee := a.actual()
	if typee.Type != SYMBOL {
		panic(fmt.Sprintf("Expectative class variable type but found: %+v", typee))
	}

	a.next()
	return VariableClass{
		Symbol: Identify{
			Val: name.Value,
		},
		Type: AssingType{
			Type: typee.Value,
		},
		Public: public,
	}
}

func (a *Ast) parseClass() Class {
	//a.next()
	a.next()
	className := a.actual()

	if className.Type != SYMBOL {
		panic(fmt.Sprintf("Expectative class name but found: %+v", className))
	}
	class := Class{
		Vars: []VariableClass{},
		Symbol: Identify{
			Val: className.Value,
		},
	}
	a.next()
	body := a.parseStmt()
	if body.Kind() != TypeBody {
		panic(fmt.Sprintf("Expectative body class but found: %+v", body))
	}
	bod := body.(BodyStatement)
	for _, v := range bod.Body {
		if v.Kind() == TypeVarDeclarationClass {
			class.Vars = append(class.Vars, v.(VariableClass))
		} else if v.Kind() == TypeFunction {
			class.Methods = append(class.Methods, v.(Function))
		} else if v.Kind() == TypeCPP {
			class.Cpps = append(class.Cpps, v.(CppCode))
		} else {
			panic(fmt.Sprintf("Invalid expression: %+v", v))

		}
	}

	return class

}
func (a *Ast) parseBody() BodyStatement {
	body := BodyStatement{
		Body: []Stmt{},
	}
	a.next()
	for {
		if a.isEOF() {
			panic("Expectative block close(\"}\") but EOF")
		}
		if a.actual().Type == CLOSE_BRACKET {
			break
		}
		body.Body = append(body.Body, a.parseStmt())

	}
	a.next()
	return body

}
