package pkg

import "runtime"

func (c *Compiler) toCppOnly(only *Only, pro *ProgramInfoCompile, scope *ScopeInfoCompile, afterPkg bool) string {
	l := ""
	for _, v := range only.OS {
		if v == WINDOWS {
			if runtime.GOOS == "windows" {
				for _, v := range only.Body.Body {

					l += c.GenCode(v, true, pro, scope, afterPkg)
				}
				return l
			}
		} else if v == LINUX {
			if runtime.GOOS == "linux" {
				for _, v := range only.Body.Body {

					l += c.GenCode(v, true, pro, scope, afterPkg)
				}
				return l
			}
		} else if v == MACOS {
			if runtime.GOOS == "darwin" {
				for _, v := range only.Body.Body {

					l += c.GenCode(v, true, pro, scope, afterPkg)
				}
				return l
			}
		}

	}
	return ""
}
