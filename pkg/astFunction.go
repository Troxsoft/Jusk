package pkg

type Argument struct {
	Symbol Identify
	Type   AssingType
}

func (a Argument) Kind() int {
	return TypeArgument
}

type Function struct {
	Symbol    Identify
	Arguments []Argument
	Body      BodyStatement
	Public    bool

	Return Identify
}
type FunctionCall struct {
	Symbol    Identify
	Arguments []Stmt
}

func (f FunctionCall) Value() any {
	return ""
}
func (f FunctionCall) Kind() int {
	return TypeFunctionCall
}
func (f Function) Kind() int {
	return TypeFunction
}
func (a Argument) Value() any {
	return ""
}
func (a *Ast) parseArgument() Argument {
	//fmt.Println(a.Tokens[2].Value)
	if a.actual().Type == SYMBOL && a.Tokens[1].Type == TWO_POINTS && a.Tokens[2].Type == CPP {
		name := Identify{
			Val: a.actual().Value.(string),
		}
		a.next()
		a.next()
		e := a.parseCpp()

		return Argument{
			Symbol: name,
			Type: AssingType{
				Type: e.Code,
			},
		}
	} else {
		if a.actual().Type != SYMBOL && a.Tokens[1].Type != TWO_POINTS && a.Tokens[2].Type != SYMBOL {
			panic("Invalid sintax on function declaration")
		}
		name := Identify{
			Val: a.actual().Value.(string),
		}
		a.next()
		e := a.parseAssingTypeExpr()

		return Argument{
			Symbol: name,
			Type:   e.(AssingType),
		}
	}

}
func (a *Ast) parseFunction() Function {
	j := a.actual()
	public := true
	arguments := []Argument{}
	if j.Type == PUBLIC {
		public = true
	} else {
		public = false
	}
	a.next()
	a.next()
	funcName := Identify{
		Val: a.actual().Value.(string),
	}
	a.next()
	a.next()
	for a.actual().Type != CLOSE_PARENT {
		if a.isEOF() {
			panic("Expectative \")\" but found EOF")
		}
		arguments = append(arguments, a.parseArgument())
	}
	a.next()

	k := a.parseStmt()
	if k.Kind() == TypeIdentify {
		k40 := a.parseBody()
		return Function{
			Symbol:    funcName,
			Arguments: arguments,
			Body:      k40,
			Public:    public,
			Return:    k.(Identify),
		}
	}
	return Function{
		Symbol:    funcName,
		Arguments: arguments,
		Body:      k.(BodyStatement),
		Public:    public,
		Return:    Identify{Val: "Void"},
	}

}

func (a *Ast) parseFunctionCall() FunctionCall {
	name := Identify{
		Val: a.actual().Value.(string),
	}
	a.next()
	p := a.parseStmt()
	//fmt.Println(p.Kind())

	//a.next()
	return FunctionCall{
		Symbol:    name,
		Arguments: p.(Parent).Children,
	}

}
