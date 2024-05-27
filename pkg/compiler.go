package pkg

import (
	"fmt"
	"strings"
)

type Compiler struct {
	Astes *Ast
	Cpp   string
}

func NewCompiler(ast *Ast) *Compiler {
	return &Compiler{
		Astes: ast,
	}
}
func (c *Compiler) preprocessType(str string) string {
	if str == "Int" {
		return "int"
	}
	return str
}
func (c *Compiler) compilePrimaryExpressions(epr Stmt) {
	if epr.Kind() == TypeLiteralNumber {
		c.Cpp += fmt.Sprint(epr.(LiteralNumeric).Value())
	} else if epr.Kind() == TypeParent {
		c.Cpp += "("
		for i, e := range epr.(Parent).Children {

			c.compilePrimaryExpressions(e)
			if i != len(epr.(Parent).Children)-1 {
				c.Cpp += ","
			}
		}
		//c.compilePrimaryExpressions(epr.(Parent).Children)
		c.Cpp += ")"
	} else if epr.Kind() == TypeIdentify {
		c.Cpp += epr.(Identify).Val.(string)
	} else if epr.Kind() == TypeAssingDeclaration {
		f := epr.(AssingDeclaration)
		c.Cpp += f.Symbol.(Identify).Val.(string)
		c.Cpp += "="
		c.compilePrimaryExpressions(f.Val.(Stmt))
		c.Cpp += ";"
	} else if epr.Kind() == TypeLiteralString {
		c.Cpp += fmt.Sprintf("\"%s\"", epr.(LiteralString).Val)
	} else if epr.Kind() == TypeVarDeclaration {
		o := epr.(VarDeclaration)

		c.Cpp += c.preprocessType(o.Type.(AssingType).Type.(string)) + " "
		c.Cpp += o.Symbol.(Identify).Val.(string)
		c.Cpp += "="

		c.compilePrimaryExpressions(o.Val.(Stmt))
		if !strings.HasSuffix(c.Cpp, ";") {
			c.Cpp += ";"

		}

	} else if epr.Kind() == TypeReturn {
		c.Cpp += "return "
		c.compilePrimaryExpressions(epr.(Return).Val.(Stmt))

		c.Cpp += ";"
	} else if epr.Kind() == TypeBody {
		pilo := epr.(BodyStatement)
		c.Cpp += "{"
		for _, e := range pilo.Body {

			c.compilePrimaryExpressions(e)
		}
		c.Cpp += "}"
	} else if epr.Kind() == TypeArgument {
		c.Cpp += epr.(Argument).Type.Type.(string) + " "
		c.Cpp += epr.(Argument).Symbol.Val.(string) + " "

	} else if epr.Kind() == TypeFunctionCall {
		o := epr.(FunctionCall)
		c.Cpp += o.Symbol.Val.(string)
		p := Parent{
			Children: o.Arguments,
		}
		c.compilePrimaryExpressions(p)
		c.Cpp += ";"
	} else if epr.Kind() == TypeFunction {
		f := epr.(Function)
		c.Cpp += c.preprocessType(f.Return.Val.(string)) + " "
		c.Cpp += f.Symbol.Val.(string) + "("
		for i, e := range f.Arguments {
			c.compilePrimaryExpressions(e)
			if i != len(f.Arguments)-1 {
				c.Cpp += ","
			}
		}
		c.Cpp += ")"
		c.compilePrimaryExpressions(f.Body)
	} else if epr.Kind() == TypeVarDeclarationClass {
		o := epr.(VariableClass)
		if o.Public {
			c.Cpp += "public: "
		} else {
			c.Cpp += "private: "
		}
		c.Cpp += fmt.Sprintf("%s %s;", c.preprocessType(o.Type.Type.(string)), o.Symbol.Val.(string))
	} else if epr.Kind() == TypeClass {
		o := epr.(Class)
		c.Cpp += "class " + o.Symbol.Val.(string) + "{"
		for _, e := range o.Cpps {
			c.compilePrimaryExpressions(e)
		}
		for _, e := range o.Vars {
			c.compilePrimaryExpressions(e)
		}
		for _, e := range o.Methods {
			c.compilePrimaryExpressions(e)
		}
		c.Cpp += "};"
	} else if epr.Kind() == TypeCPP {
		c.Cpp += epr.(CppCode).Code
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
	} else {
		panic(epr.Kind())
	}
}
func (c *Compiler) Compile() {
	program := c.Astes.Nodes.(ProgramBody)
	for _, e := range program.Body {
		c.compilePrimaryExpressions(e)
	}
}
