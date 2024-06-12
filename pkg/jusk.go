package pkg

import "strings"

const (
	VERSION = "0.1.0-experimental"
)

type Jusk struct {
	Code     string
	Tokens   []Token
	Astes    *Ast
	Compiles *Compiler
}

func (c *Jusk) Compile(path string, father *Compiler) string {
	r := NewCompiler(c.Astes, path, father)
	r.Compile()
	c.Compiles = r

	return r.Cpp
}
func NewJuskLang(code string) *Jusk {
	code = strings.TrimSpace(code)
	code = strings.ReplaceAll(code, "\r\n", "\n")
	code += "\n\n"
	return &Jusk{
		Code:   code,
		Tokens: []Token{},
		Astes:  nil,
	}
}
