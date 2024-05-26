package pkg

import "strings"

type Jusk struct {
	Code   string
	Tokens []Token
	Astes  *Ast
}

func (c *Jusk) Compile() string {
	r := NewCompiler(c.Astes)
	r.Compile()
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
