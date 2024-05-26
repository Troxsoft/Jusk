package pkg

import (
	"fmt"
	"strconv"
)

func (j *Jusk) Tokenize() error {
	i := 0
	toks := []Token{}
	line := 0
	for i < len(j.Code) {
		ac := j.Code[i]
		if ac == '\n' {
			//toks = append(toks, NewToken(ENDLINE, "endline"))
			line++
			i++
		} else if ac == '"' {

			start := i + 1
			end := i + 1
			for (end < len(j.Code) && j.Code[end] != '"') || (end < len(j.Code) && j.Code[end] != '"' && j.Code[end-1] != '\\') {
				//fmt.Println(string(j.Code[end]))
				if end+1 < len(j.Code) {
					end++
				} else {
					end++
					panic("Expectative close string but found:EOF")
					//break
				}
			}
			toks = append(toks, NewToken(STRING, j.Code[start:end]))
			i = end + 1

		} else if i+3 < len(j.Code) && ac == 'v' && j.Code[i+1] == 'a' && j.Code[i+2] == 'r' && j.Code[i+3] == ' ' {
			toks = append(toks, NewToken(VAR, "var"))
			i += 3
		} else if (ac >= 'A' && ac <= 'Z') || (ac >= 'a' && ac <= 'z') || ac == '_' {
			start := i
			end := i
			for end < len(j.Code) && isIntNumber(j.Code[end]) || (j.Code[end] >= 'A' && j.Code[end] <= 'Z') || (j.Code[end] >= 'a' && j.Code[end] <= 'z') || (j.Code[end] == '_') {
				if end+1 < len(j.Code) {

					end++
				} else {
					end++
					break
				}
			}
			toks = append(toks, NewToken(SYMBOL, j.Code[start:end]))
			i = end
		} else if ac == ' ' {
			//toks = append(toks, NewToken(SPACE, "space"))
			i++
		} else if ac == '+' {
			toks = append(toks, NewToken(PLUS, "+"))
			i++
		} else if ac == '%' {
			toks = append(toks, NewToken(PORCENT, "%"))
			i++
		} else if ac == '-' {
			toks = append(toks, NewToken(MINUS, "-"))
			i++
		} else if ac == '/' {
			toks = append(toks, NewToken(DIVIDE, "/"))
			i++
		} else if ac == '=' {
			toks = append(toks, NewToken(EQUAL, "="))
			i++
		} else if ac == '*' {
			toks = append(toks, NewToken(MULTIPLY, "*"))
			i++
		} else if ac == '(' {
			toks = append(toks, NewToken(OPEN_PARENT, "("))
			i++
		} else if ac == ')' {
			toks = append(toks, NewToken(CLOSE_PARENT, ")"))
			i++
		} else if ac == ':' {
			toks = append(toks, NewToken(TWO_POINTS, ":"))
			i++
		} else if i+1 < len(j.Code) && ac == '=' && j.Code[i+1] == '=' {
			toks = append(toks, NewToken(TWO_POINTS, ":"))
			i++
		} else if isIntNumber(ac) {
			start := i
			end := i
			floatb := false
			for end < len(j.Code) && isIntNumber(j.Code[end]) {
				if end+1 < len(j.Code) && j.Code[end+1] == '.' && !floatb {

					floatb = true
					end++
				}
				end++
			}

			if floatb {
				num, _ := strconv.ParseFloat(j.Code[start:end], 64)
				toks = append(toks, NewToken(FLOAT, num))
			} else {
				num, _ := strconv.Atoi(j.Code[start:end])
				toks = append(toks, NewToken(INT, num))
			}
			i = end
		} else {
			return fmt.Errorf("invalid word: %s at (%v,%v)", string(ac), line, i)
		}
	}
	toks = append(toks, NewToken(EOF, "end of code"))
	j.Tokens = toks

	return nil
}
