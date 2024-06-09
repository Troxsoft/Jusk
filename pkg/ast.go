package pkg

import (
	"fmt"
)

const (
	TypeBinaryExpression = iota
	TypeLiteralNumber
	TypeLiteralString
	TypeBodyStatementFile
	TypeIdentify
	TypeAssingDeclaration
	TypeAsingType
	TypeVarDeclaration
	TypeVarDeclarationStruct
	TypeParent
	TypeBody
	TypeStruct
	TypeCPP
	TypeArgument
	TypeFunction
	TypeReturn
	TypeFunctionCall
	TypeImport
	TypePointStmt
	TypeBoolean
	TypeIf
)

type LiteralNumeric struct {
	Val     any
	IsFloat bool
}
type LiteralString struct {
	Val string
}
type Return struct {
	Val any
}
type Import struct {
	Val any
}

func (c Import) Kind() int {
	return TypeImport
}
func (c Import) Value() any {
	return c.Val
}
func (l Return) Value() any {
	return l.Val
}
func (l Return) Kind() int {
	return TypeReturn
}

type Parent struct {
	Children []Stmt
}

func (l Parent) Value() any {
	return ""
}
func (l Parent) Kind() int {
	return TypeParent
}
func (l LiteralString) Value() any {
	return l.Val
}
func (l LiteralString) Kind() int {
	return TypeLiteralString
}

type VarDeclaration struct {
	Symbol any
	Type   any
	Val    any
}

func (v VarDeclaration) Kind() int {
	return TypeVarDeclaration
}
func (v VarDeclaration) Value() any {
	return v.Val
}

func (l LiteralNumeric) Value() any {
	return l.Val
}
func (l LiteralNumeric) Kind() int {
	return TypeLiteralNumber
}

type BinaryExpression struct {
	Left     Expr
	Operator int
	Right    Expr
}
type PointStmt struct {
	Children any
	Father   any
}

func (s PointStmt) Kind() int {
	return TypePointStmt
}
func (s PointStmt) Value() any {
	return nil
}

type Identify struct {
	Val any
}

func (l Identify) Value() any {
	return l.Val
}
func (l Identify) Kind() int {
	return TypeIdentify
}

// Kind implements Expr.
func (l BinaryExpression) Kind() int {
	return TypeBinaryExpression
}

// Value implements Expr.
func (l BinaryExpression) Value() any {
	panic("unimplemented")
}

type ProgramBody struct {
	Body    []Stmt
	PkgName string
}
type AssingType struct {
	Type any
}

// Kind implements Expr.
func (l AssingType) Kind() int {
	return TypeAsingType
}

// Value implements Expr.
func (l AssingType) Value() any {
	return l.Type
}

type AssingDeclaration struct {
	Symbol any
	Val    any
}

// Kind implements Expr.
func (l AssingDeclaration) Kind() int {
	return TypeAssingDeclaration
}

// Value implements Expr.
func (l AssingDeclaration) Value() any {
	return l.Val
}

func (l ProgramBody) Kind() int {
	return TypeBodyStatementFile
}

type Ast struct {
	Nodes  Stmt
	Tokens []Token
}
type Expr interface {
	Value() any
	Kind() int
}
type Stmt interface {
	Kind() int
}

func NewAst(toks []Token) *Ast {
	return &Ast{
		Tokens: toks,
	}
}
func (a *Ast) ProduceAst() {
	program := ProgramBody{
		Body: []Stmt{},
	}
	if a.actual().Type != PACKAGE && a.Tokens[1].Type != SYMBOL {
		panic(fmt.Sprintf("Expectative package name but found: %+v , %+v", a.actual(), a.Tokens[1]))
	}
	name := a.Tokens[1].Value.(string)
	program.PkgName = name
	a.next()
	a.next()
	for !a.isEOF() {

		program.Body = append(program.Body, a.parseStmt(false))
	}
	a.Nodes = program

}
func (a *Ast) parseVarDeclaration() Stmt {
	a.next()
	VarName := a.actual()
	if VarName.Type != SYMBOL {
		panic("Expectative symbol")
	}
	f := a.parseStmt(false)
	//a.next()
	//fmt.Printf("%+v\n", f)

	typeVar := a.parseAssingTypeExpr()
	//fmt.Printf("%+v\n", typeVar)
	a.next()
	assing := a.parseStmt(false)
	//a.next()
	return VarDeclaration{
		Symbol: f,
		Type:   typeVar,
		Val:    assing,
	}

}
func (a *Ast) parseStmt(h bool) Stmt {
	typeee := a.actual().Type
	//fmt.Println(typeee)
	if typeee == VAR {
		return a.parseVarDeclaration()
	} else if typeee == IF {
		return a.parseIf()
	} else if (typeee == PUBLIC || typeee == PRIVATE) && a.Tokens[1].Type == SYMBOL {
		vari := a.parseVariableClass()
		return vari
	} else if typeee == IMPORT && a.Tokens[1].Type == STRING {
		pedro := a.Tokens[1]
		a.next()
		a.next()
		return Import{
			Val: pedro.Value,
		}
	} else if typeee == OPEN_BRACKET {

		body := a.parseBody()
		return body
	} else if typeee == STRUCT {
		class := a.parseStruct()
		return class
	} else if (typeee == PUBLIC || typeee == PRIVATE) && a.Tokens[1].Type == FUNCTION {
		vari := a.parseFunction()
		return vari
	} else {
		return a.parseExpr()

	}
}

//return a.parseExpr()

func (a *Ast) parseExpr() Expr {
	typeee := a.actual().Type
	if typeee == SYMBOL && a.Tokens[1].Type == POINT {
		//fmt.Printf("jnjunj %+v lol %+v \n ", a.actual(), a.Tokens[1])
		a1 := a.actual()
		a.next()
		a.next()
		f := a.parseStmt(false)

		return PointStmt{
			Children: f,
			Father:   a1.Value,
		}
	} else if typeee == SYMBOL && a.Tokens[1].Type == OPEN_PARENT {

		return a.parseFunctionCall()
	} else if typeee == CPP {
		cpp := a.parseCpp()
		return cpp
	} else {

		return a.parseAssingExpr()

	}

}
func (a *Ast) parseAAAFuncCall() Expr {
	if a.actual().Type == SYMBOL && a.Tokens[1].Type == OPEN_PARENT {
		return a.parseFunctionCall()
	} else {
		return a.parsePrimaryExpr()
	}
}
func (a *Ast) parseAAAPointStmtAsExpr() Expr {
	if a.actual().Type == SYMBOL && a.Tokens[1].Type == POINT {
		a1 := a.actual()
		a.next()
		a.next()
		f := a.parseExpr()

		//a.Tokens = append([]Token{},,a.Tokens... )
		return PointStmt{
			Children: f,
			Father:   a1.Value,
		}
	} else {
		return a.parseAAAFuncCall()
	}
}
func (a *Ast) parseAssingExpr() Expr {
	left := a.parseAdditiveExpr()
	for a.actual().Type == EQUAL {

		//Eq := a.actual()

		a.next()
		right := a.parseStmt(false)
		left = AssingDeclaration{
			Symbol: left,
			Val:    right,
		}

	}
	return left
}
func (a *Ast) parseAdditiveExpr() Expr {
	// + -
	left := a.parseMultiplyExpr()
	for a.actual().Type == PLUS || a.actual().Type == MINUS {
		operator := a.actual()
		a.next()
		right := a.parseExpr()
		left = BinaryExpression{
			Left:     left,
			Operator: operator.Type,
			Right:    right,
		}

	}
	return left
}

func (a *Ast) parseLogicExpr() Expr {
	left := a.parseAAAPointStmtAsExpr()

	// * / %
	for a.actual().Type == COMPARE || a.actual().Type == LESS || a.actual().Type == GREATER || a.actual().Type == COMPARE_GREATER || a.actual().Type == COMPARE_LESS || a.actual().Type == NOCOMPARE {
		operator := a.actual()
		a.next()
		right := a.parseExpr()

		left = BinaryExpression{
			Left:     left.(Expr),
			Operator: operator.Type,
			Right:    right.(Expr),
		}

	}
	return left.(Expr)
}
func (a *Ast) parseMultiplyExpr() Expr {
	// 3
	left := a.parseLogicExpr()
	for a.actual().Type == MULTIPLY || a.actual().Type == DIVIDE || a.actual().Type == PORCENT {
		operator := a.actual()
		a.next()
		right := a.parseExpr()
		left = BinaryExpression{
			Left:     left,
			Operator: operator.Type,
			Right:    right,
		}

	}
	return left
}
func (a *Ast) parseAssingTypeExpr() Expr {

	//fmt.Println("jnjn ", q)
	//fmt.Println("ekoe ", a.actual())
	if a.actual().Type == TWO_POINTS {

		a.next()

		d := a.actual()
		if d.Type != SYMBOL {
			panic(fmt.Sprintf("Expectative SYMBOL but found: %+v", d))
		}
		a.next()
		left := AssingType{
			Type: d.Value,
		}
		return left
	} else {
		return a.parseExpr()
	}
}
func (a *Ast) parsePrimaryExpr() Expr {
	tok := a.actual()
	switch tok.Type {
	case BOOLEAN:
		{
			a.next()
			return Boolean{
				Bool: tok.Value.(bool),
			}
		}
	case INT:
		{
			a.next()
			return LiteralNumeric{
				Val:     tok.Value,
				IsFloat: false,
			}
		}
	case FLOAT:
		{
			a.next()
			return LiteralNumeric{
				Val:     tok.Value,
				IsFloat: true,
			}
		}
	case SYMBOL:
		{

			a.next()
			return Identify{
				Val: tok.Value,
			}

		}
	case OPEN_PARENT:
		{
			ds := []Stmt{}
			a.next()
			if a.actual().Type == CLOSE_PARENT {
				a.next()
				return Parent{
					Children: ds,
				}
			}
			for {
				e := a.parseStmt(false)
				if a.actual().Type != CLOSE_PARENT {
					if a.actual().Type == COMMA {

						ds = append(ds, e)
						a.next()
					}
				} else {
					ds = append(ds, e)
					a.next()
					break
				}

				//g := a.actual()
				//a.next()
			}

			return Parent{
				Children: ds,
			}

		}
	case EOF:
		{
			return Identify{
				"",
			}
		}
	case STRING:
		{
			a.next()
			return LiteralString{
				Val: tok.Value.(string),
			}
		}
	case RETURN:
		{
			a.next()
			return Return{
				Val: a.parseStmt(false),
			}
		}
	default:
		{

			//log.Fatalf("Invalid Token: %+v\n", tok)
			panic(fmt.Sprintf("Invalid token %+v", tok))
		}
	}

}
func (a *Ast) isEOF() bool {
	return a.Tokens[0].Type == EOF
}
func (a *Ast) actual() Token {
	return a.Tokens[0]
}
func (a *Ast) next() {
	a.Tokens = a.Tokens[1:]
}
