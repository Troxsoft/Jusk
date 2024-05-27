package pkg

const (
	INT = iota
	FLOAT
	SYMBOL
	SPACE
	STRING
	ENDLINE
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	EOF
	PORCENT
	OPEN_PARENT
	CLOSE_PARENT
	OPEN_BRACKET
	CLOSE_BRACKET
	VAR
	TWO_POINTS
	EQUAL
	PRIVATE
	PUBLIC
	CLASS
	CPP
	FUNCTION
	RETURN
	COMMA
	PACKAGE
)

func isValidSize(s string, index, l int) bool {
	if index+l < len(s) {
		return true
	}
	return false
}

type Token struct {
	Type  int
	Value any
}

func NewToken(t int, v any) Token {
	return Token{
		Type:  t,
		Value: v,
	}
}
