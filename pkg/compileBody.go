package pkg

func (c *Compiler) toCppBody(b BodyStatement, pro *ProgramInfoCompile, scope *ScopeInfoCompile, afterPkg bool) (*ScopeInfoCompile, string) {
	p := "{"
	scopeNew := &ScopeInfoCompile{
		vars: scope.vars,
	}
	for _, v := range b.Body {

		p += c.GenCode(v, true, pro, scopeNew, afterPkg)
	}

	return scopeNew, p + "}"
}
