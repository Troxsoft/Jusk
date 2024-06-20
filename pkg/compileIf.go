package pkg

import "fmt"

func (c *Compiler) parseIf(ifC IFCondition, pro *ProgramInfoCompile, scope *ScopeInfoCompile, afterPkg bool) string {

	if ifC.OP != nil {
		k := ifC.OP
		l := fmt.Sprintf("if (%s)%s", c.GenCode(*k, true, pro, scope, afterPkg), c.GenCode(ifC.Body, false, pro, scope, afterPkg))
		for _, v := range ifC.Elifs {
			if v.OP != nil {
				l += fmt.Sprintf("else if (%s)%s", c.GenCode(*v.OP, true, pro, scope, afterPkg), c.GenCode(v.Body, false, pro, scope, afterPkg))
			} else {
				if v.Bool.Bool {

					l += fmt.Sprintf("else if (true)%s", c.GenCode(v.Body, false, pro, scope, afterPkg))
				} else {
					l += fmt.Sprintf("else if (false)%s", c.GenCode(v.Body, false, pro, scope, afterPkg))

				}
			}
		}
		if ifC.Else != nil {
			l += "else " + c.GenCode(ifC.Else.Body, false, pro, scope, afterPkg)
		}
		return l
	} else {
		if ifC.Bool.Bool {
			l := fmt.Sprintf("if (true)%s", c.GenCode(ifC.Body, false, pro, scope, afterPkg))
			for _, v := range ifC.Elifs {
				if v.OP != nil {
					l += fmt.Sprintf("else if (%s)%s", c.GenCode(*v.OP, true, pro, scope, afterPkg), c.GenCode(v.Body, false, pro, scope, afterPkg))
				} else {
					if v.Bool.Bool {

						l += fmt.Sprintf("else if (true)%s", c.GenCode(v.Body, false, pro, scope, afterPkg))
					} else {
						l += fmt.Sprintf("else if (false)%s", c.GenCode(v.Body, false, pro, scope, afterPkg))

					}
				}
			}
			if ifC.Else != nil {
				l += "else " + c.GenCode(ifC.Else.Body, false, pro, scope, afterPkg)
			}
			return l
		} else {
			l := fmt.Sprintf("if (false)%s", c.GenCode(ifC.Body, false, pro, scope, afterPkg))
			for _, v := range ifC.Elifs {
				if v.OP != nil {
					l += fmt.Sprintf("else if (%s)%s", c.GenCode(*v.OP, true, pro, scope, afterPkg), c.GenCode(v.Body, false, pro, scope, afterPkg))
				} else {
					if v.Bool.Bool {

						l += fmt.Sprintf("else if (true)%s", c.GenCode(v.Body, false, pro, scope, afterPkg))
					} else {
						l += fmt.Sprintf("else if (false)%s", c.GenCode(v.Body, false, pro, scope, afterPkg))

					}
				}
			}
			if ifC.Else != nil {
				l += "else " + c.GenCode(ifC.Else.Body, false, pro, scope, afterPkg)
			}
			return l
		}
	}

}
