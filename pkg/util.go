package pkg

func isIntNumber(num byte) bool {
	if num == '0' || num == '1' || num == '2' || num == '3' || num == '4' || num == '5' || num == '6' || num == '7' || num == '8' || num == '9' {
		return true
	}
	return false
}

func replaceTypesPrimitivesForCppType(c string) string {
	if c == "Int" {
		return "int"
	} else if c == "Void" {
		return "void"
	} else if c == "Float" {
		return "float"
	} else if c == "Bool" {
		return "bool"
	} else if c == "Str" {
		return "std::string"
	} else {
		return c
	}
}
