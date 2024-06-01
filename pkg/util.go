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
	} else {
		return c
	}
}
