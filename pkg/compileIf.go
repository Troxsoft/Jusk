package pkg

import "fmt"

func (c *Compiler) parseIf(ifC IFCondition, pro *ProgramInfoCompile, scope *ScopeInfoCompile, afterPkg bool) string {

	if ifC.OP != nil {
		k := ifC.OP
		return fmt.Sprintf("if (%s)%s", c.GenCode(*k, true, pro, scope, afterPkg), c.GenCode(ifC.Body, false, pro, scope, afterPkg))
	} else {
		if ifC.Bool.Bool {
			return fmt.Sprintf("if (true)%s", c.GenCode(ifC.Body, false, pro, scope, afterPkg))
		} else {
			return fmt.Sprintf("if (false)%s", c.GenCode(ifC.Body, false, pro, scope, afterPkg))

		}
	}

}
