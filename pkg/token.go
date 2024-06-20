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
	POINT
	TWO_POINTS
	EQUAL
	PRIVATE
	PUBLIC
	STRUCT
	CPP
	FUNCTION
	RETURN
	COMMA
	PACKAGE
	IMPORT
	COMMENT
	IF
	ONLY
	WINDOWS
	MACOS
	LINUX
	NEWVAR // :=
	// logic op
	COMPARE         // == igual
	NOCOMPARE       // != no igual
	GREATER         // > mayor
	LESS            // < menor
	COMPARE_GREATER // >= mayor o igual
	BOOLEAN         // true or false
	TYPE            // @type
	COMPARE_LESS    // <= menor o igual
	ELIF            // elif
	ElSE            // else
	AND             // and
	OR              // or

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
