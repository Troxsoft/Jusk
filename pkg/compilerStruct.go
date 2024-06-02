package pkg

import "fmt"

func (c *Compiler) getStruct(name string, pro *ProgramInfoCompile) *StructInfoCompile {
	for _, v := range pro.structs {
		if v.Name == name {
			return v
		}
	}
	return nil
}
func (c *Compiler) toCppStruct(s StructInfoCompile, pro *ProgramInfoCompile, scope *ScopeInfoCompile) string {
	l := "struct "
	l += s.Name + " "
	for _, v := range s.Vars {
		scope.vars = append(scope.vars, &VarInfoCompile{
			Type: v.Type,
			Name: v.Name,
		})
	}
	l += "{"
	for _, v := range s.Cpps {

		l += v.Code
	}
	for _, v := range s.Vars {
		if v.Public {
			l += "public: "
		} else {
			l += "private: "
		}
		l += fmt.Sprintf("%s %s;", replaceTypesPrimitivesForCppType(v.Type), v.Name)
	}
	l += "};"

	return l
}
