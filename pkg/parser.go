package pkg

func (s *Jusk) GenerateAst() error {
	ast := NewAst(s.Tokens)

	ast.ProduceAst()

	s.Astes = ast

	return nil
}
