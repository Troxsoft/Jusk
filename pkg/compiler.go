package pkg

import "fmt"

type Compiler struct {
	Astes *Ast
	Cpp   string
}

func NewCompiler(ast *Ast) *Compiler {
	return &Compiler{
		Astes: ast,
	}
}

func (c *Compiler) compilePrimaryExpressions(epr Stmt) {
	if epr.Kind() == TypeLiteralNumber {
		c.Cpp += fmt.Sprint(epr.(LiteralNumeric).Value())
	} else if epr.Kind() == TypeParent {
		c.Cpp += "("
		c.compilePrimaryExpressions(epr.(Parent).Children)
		c.Cpp += ")"
	} else if epr.Kind() == TypeIdentify {
		c.Cpp += epr.(Identify).Val.(string)
	} else if epr.Kind() == TypeAssingDeclaration {
		f := epr.(AssingDeclaration)
		c.Cpp += f.Symbol.(Identify).Val.(string)
		c.Cpp += "="
		c.compilePrimaryExpressions(f.Val.(Stmt))
		c.Cpp += ";"
	}else if epr.Kind() == TypeLiteralString{
		c.Cpp += fmt.Sprintf("\"%s\"",epr.(LiteralString).Val)
	}else if epr.Kind() == TypeVarDeclaration{
		o := epr.(VarDeclaration)
		c.Cpp += o.Type.(AssingType).Type.(string)+" "
		c.Cpp += o.Symbol.(Identify).Val.(string)
		c.Cpp += "="
		c.compilePrimaryExpressions(o.Val.(Stmt))
		c.Cpp += ";"

	} else if epr.Kind() == TypeBinaryExpression {
		f := epr.(BinaryExpression)
		if f.Operator == PLUS {
			c.compilePrimaryExpressions(f.Left)
			c.Cpp += " + "
			c.compilePrimaryExpressions(f.Right)
		} else if f.Operator == MINUS {
			c.compilePrimaryExpressions(f.Left)
			c.Cpp += " - "
			c.compilePrimaryExpressions(f.Right)
		} else if f.Operator == MULTIPLY {
			c.compilePrimaryExpressions(f.Left)
			c.Cpp += " * "
			c.compilePrimaryExpressions(f.Right)
		} else if f.Operator == DIVIDE {
			c.compilePrimaryExpressions(f.Left)
			c.Cpp += " / "
			c.compilePrimaryExpressions(f.Right)
		} else if f.Operator == PORCENT {
			c.compilePrimaryExpressions(f.Left)
			c.Cpp += " % "
			c.compilePrimaryExpressions(f.Right)
		}
	}
}
func (c *Compiler) Compile() {
	program := c.Astes.Nodes.(ProgramBody)
	for _, e := range program.Body {
		c.compilePrimaryExpressions(e)
	}
}
