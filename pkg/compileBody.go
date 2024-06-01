package pkg

func (c *Compiler) toCppBody(b BodyStatement, pro *ProgramInfoCompile, scope *ScopeInfoCompile) (*ScopeInfoCompile, string) {
	p := "{"
	scopeNew := &ScopeInfoCompile{
		vars: scope.vars,
	}
	for _, v := range b.Body {

		p += c.GenCode(v, pro, scopeNew)
	}

	return scopeNew, p + "}"
}
