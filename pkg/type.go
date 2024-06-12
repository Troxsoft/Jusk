package pkg

import (
	"fmt"
)

func (c *Compiler) getType(value Expr, scope *ScopeInfoCompile, pro *ProgramInfoCompile) (*Identify, error) {

	if value.Kind() == TypeLiteralNumber {
		o := value.(LiteralNumeric)
		if o.IsFloat {
			return &Identify{Val: "Float"}, nil
		} else {
			return &Identify{Val: "Int"}, nil
		}
	} else if value.Kind() == TypeType {
		return &Identify{Val: "Str"}, nil
	} else if value.Kind() == TypePointStmt {
		o := value.(PointStmt)
		var pkg *Compiler = nil
		if c.Father != nil {
			for _, v := range c.Father.Imports {
				if o.Father.(string) == v.Name {
					pkg = v.Compilers
				}
			}
		}
		for _, v := range c.Imports {
			if o.Father.(string) == v.Name {
				pkg = v.Compilers
			}
		}
		if pkg == nil {
			panic(fmt.Sprintf("Unknown SYMBOL %s", o.Father.(string)))
		}
		p, err := pkg.getType(o.Children.(Expr), scope, pro)
		if err != nil {
			return nil, err
		}
		return p, nil
	} else if value.Kind() == TypeFunctionCall {
		o := value.(FunctionCall)

		var fun *FunctionInfoCompile = nil

		for _, v := range pro.funcs {
			if o.Symbol.Val.(string) == v.Name {
				fun = v
			}
		}
		for _, v := range c.Public.funcs {
			if o.Symbol.Val.(string) == v.Name {
				fun = v
			}
		}
		if fun == nil {
			return nil, fmt.Errorf("Not exists Function: %s", o.Symbol.Val.(string))
		}
		return &Identify{Val: fun.ReturnType}, nil

	} else if value.Kind() == TypeLiteralString {
		return &Identify{Val: "Str"}, nil
	} else if value.Kind() == TypeIdentify {
		o := value.(Identify)
		p := c.getVar(o.Val.(string), scope, pro)
		if p == nil {
			return nil, fmt.Errorf("Unknown SYMBOL: %s", p.Name)
		}
		return &Identify{Val: p.Type}, nil
	} else if value.Kind() == TypeParent {
		o := value.(Parent)
		if len(o.Children) == 0 {
			return nil, fmt.Errorf("Invalid type: %+v", value)
		}
		return c.getType(o.Children[0].(Expr), scope, pro)
	} else if value.Kind() == TypeBinaryExpression {
		o := value.(BinaryExpression)
		boolean := false
		if o.Operator == COMPARE || o.Operator == LESS || o.Operator == GREATER || o.Operator == COMPARE_GREATER || o.Operator == COMPARE_LESS || o.Operator == NOCOMPARE {
			boolean = true
		}
		leftType, err := c.getType(o.Left, scope, pro)
		if err != nil {
			return nil, fmt.Errorf("Error ocurred into left for binaryExpression: %+v \n%+v", o, err)
		}
		rightType, err := c.getType(o.Right, scope, pro)
		if err != nil {
			return nil, fmt.Errorf("Error ocurred into right for binaryExpression: %+v \n%+v", o, err)
		}
		if boolean {
			return &Identify{Val: "Bool"}, nil
		} else {
			if leftType.Val != rightType.Val {
				return nil, fmt.Errorf("Expectative type: %s,but found: %s", leftType.Val, rightType.Val)
			} else {
				return leftType, nil
			}
		}

	}
	return nil, fmt.Errorf("Invalid type: %+v", value)
}
