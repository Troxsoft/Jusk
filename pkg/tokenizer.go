package pkg

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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
		} else if i+5 < len(j.Code) && ac == '@' && j.Code[i+1] == 'o' && j.Code[i+2] == 'n' && j.Code[i+3] == 'l' && j.Code[i+4] == 'y' && j.Code[i+5] == ':' {
			toks = append(toks, NewToken(ONLY, "@only:"))
			i += 6 // windows
		} else if i+5 < len(j.Code) && ac == 'l' && j.Code[i+1] == 'i' && j.Code[i+2] == 'n' && j.Code[i+3] == 'u' && j.Code[i+4] == 'x' && j.Code[i+5] == ' ' {
			toks = append(toks, NewToken(LINUX, "linux"))
			i += 6 // windows
		} else if i+5 < len(j.Code) && ac == 'm' && j.Code[i+1] == 'a' && j.Code[i+2] == 'c' && j.Code[i+3] == 'o' && j.Code[i+4] == 's' && j.Code[i+5] == ' ' {
			toks = append(toks, NewToken(MACOS, "macos"))
			i += 6 // windows
		} else if i+6 < len(j.Code) && ac == 'w' && j.Code[i+1] == 'i' && j.Code[i+2] == 'n' && j.Code[i+3] == 'd' && j.Code[i+4] == 'o' && j.Code[i+5] == 'w' && j.Code[i+6] == 's' && j.Code[i+7] == ' ' {
			toks = append(toks, NewToken(WINDOWS, "windows"))
			i += 7
		} else if i+4 < len(j.Code) && ac == '@' && j.Code[i+1] == 't' && j.Code[i+2] == 'y' && j.Code[i+3] == 'p' && j.Code[i+4] == 'e' {
			toks = append(toks, NewToken(TYPE, "@type"))
			i += 5
		} else if i+4 < len(j.Code) && ac == 'e' && j.Code[i+1] == 'l' && j.Code[i+2] == 'i' && j.Code[i+3] == 'f' && j.Code[i+4] == ' ' {
			toks = append(toks, NewToken(ELIF, "elif"))
			i += 4
		} else if i+4 < len(j.Code) && ac == 'e' && j.Code[i+1] == 'l' && j.Code[i+2] == 's' && j.Code[i+3] == 'e' && j.Code[i+4] == ' ' {
			toks = append(toks, NewToken(ElSE, "else"))
			i += 4
		} else if i+1 < len(j.Code) && ac == '<' && j.Code[i+1] == '=' {
			toks = append(toks, NewToken(COMPARE_LESS, "<="))
			i += 2
		} else if i+1 < len(j.Code) && ac == 'o' && j.Code[i+1] == 'r' {
			toks = append(toks, NewToken(OR, "or"))
			i += 2
		} else if i+1 < len(j.Code) && ac == 'a' && j.Code[i+1] == 'n' && j.Code[i+2] == 'd' {
			toks = append(toks, NewToken(AND, "and"))
			i += 3
		} else if i+1 < len(j.Code) && ac == ':' && j.Code[i+1] == '=' {
			toks = append(toks, NewToken(NEWVAR, ":="))
			i += 2
		} else if i+1 < len(j.Code) && ac == '!' && j.Code[i+1] == '=' {
			toks = append(toks, NewToken(NOCOMPARE, "!="))
			i += 2
		} else if i+4 < len(j.Code) && ac == 't' && j.Code[i+1] == 'r' && j.Code[i+2] == 'u' && j.Code[i+3] == 'e' && j.Code[i+4] == ' ' {
			toks = append(toks, NewToken(BOOLEAN, true))
			i += 5
		} else if i+4 < len(j.Code) && ac == 'f' && j.Code[i+1] == 'a' && j.Code[i+2] == 'l' && j.Code[i+3] == 's' && j.Code[i+4] == 'e' && j.Code[i+5] == ' ' {
			toks = append(toks, NewToken(BOOLEAN, false))
			i += 6
		} else if i+4 < len(j.Code) && ac == 't' && j.Code[i+1] == 'r' && j.Code[i+2] == 'u' && j.Code[i+3] == 'e' && j.Code[i+4] == '\n' {
			toks = append(toks, NewToken(BOOLEAN, true))
			i += 5
		} else if i+4 < len(j.Code) && ac == 'f' && j.Code[i+1] == 'a' && j.Code[i+2] == 'l' && j.Code[i+3] == 's' && j.Code[i+4] == 'e' && j.Code[i+5] == '\n' {
			toks = append(toks, NewToken(BOOLEAN, false))
			i += 6
		} else if i+4 < len(j.Code) && ac == 't' && j.Code[i+1] == 'r' && j.Code[i+2] == 'u' && j.Code[i+3] == 'e' && j.Code[i+4] == ')' {
			toks = append(toks, NewToken(BOOLEAN, true))
			i += 4
		} else if i+4 < len(j.Code) && ac == 'f' && j.Code[i+1] == 'a' && j.Code[i+2] == 'l' && j.Code[i+3] == 's' && j.Code[i+4] == 'e' && j.Code[i+5] == ')' {
			toks = append(toks, NewToken(BOOLEAN, false))
			i += 5
		} else if i+1 < len(j.Code) && ac == '>' && j.Code[i+1] == '=' {
			toks = append(toks, NewToken(COMPARE_GREATER, ">="))
			i += 2
		} else if i+1 < len(j.Code) && ac == '=' && j.Code[i+1] == '=' {
			toks = append(toks, NewToken(COMPARE, "=="))
			i += 2
		} else if i+1 < len(j.Code) && ac == 'i' && j.Code[i+1] == 'f' && j.Code[i+2] == ' ' {
			toks = append(toks, NewToken(IF, "if"))
			i += 3
		} else if ac == '<' {
			toks = append(toks, NewToken(LESS, "<"))
			i++
		} else if ac == '>' {
			toks = append(toks, NewToken(GREATER, ">"))
			i++
		} else if ac == '"' {

			start := i + 1
			end := i + 1
			for (end < len(j.Code) && j.Code[end] != '"') || (end < len(j.Code) && j.Code[end] != '"' && string(j.Code[end-1]) != `\`) {

				if end+1 < len(j.Code) {

					if end+1 < len(j.Code) && j.Code[end] == '\\' && j.Code[end+1] == '"' {
						end++
						end++
					} else {

						end++
					}
				} else {
					end++
					panic("Expectative close string but found:EOF")
					//break
				}
			}
			//str = strings.ReplaceAll(str, `\`, "\\")
			r00 := time.Now().String()
			str := strings.ReplaceAll(j.Code[start:end], `\"`, "\"")
			str = strings.ReplaceAll(str, `\\n`, `@#@@sq`+r00)

			str = strings.ReplaceAll(str, `\n`, "\n")
			str = strings.ReplaceAll(str, `@#@@sq`+r00, "\\n")
			toks = append(toks, NewToken(STRING, str))
			i = end + 1

		} else if ac == '`' {

			start := i + 1
			end := i + 1
			for (end < len(j.Code) && j.Code[end] != '`') || (end < len(j.Code) && j.Code[end] != '`' && string(j.Code[end-1]) != `\`) {

				if end+1 < len(j.Code) {

					if end+1 < len(j.Code) && j.Code[end] == '\\' && j.Code[end+1] == '`' {
						end++
						end++
					} else {

						end++
					}
				} else {
					end++
					panic("Expectative close string but found:EOF")
					//break
				}
			}
			r00 := time.Now().String()
			str := strings.ReplaceAll(j.Code[start:end], "\\`", "`")
			str = strings.ReplaceAll(str, `\\n`, `@#@@sq`+r00)

			str = strings.ReplaceAll(str, `\n`, "\n")
			str = strings.ReplaceAll(str, `@#@@sq`+r00, "\\n")

			toks = append(toks, NewToken(STRING, str))
			i = end + 1

		} else if i+1 < len(j.Code) && ac == '/' && j.Code[i+1] == '/' {
			startComment := i
			endComment := startComment
			for endComment < len(j.Code) && j.Code[endComment] != '\n' {
				//fmt.Printf("'%s'\n", string(j.Code[endComment]))
				endComment++
			}

			i = endComment

		} else if i+1 < len(j.Code) && ac == '/' && j.Code[i+1] == '*' {
			startComment := i
			endComment := startComment
			for endComment < len(j.Code) {
				//fmt.Printf("'%s'\n", string(j.Code[endComment]))
				if j.Code[endComment] == '*' && j.Code[endComment+1] == '/' {
					endComment++
					endComment++

					break
				}
				endComment++
			}

			i = endComment

		} else if i+6 < len(j.Code) && ac == 'r' && j.Code[i+1] == 'e' && j.Code[i+2] == 't' && j.Code[i+3] == 'u' && j.Code[i+4] == 'r' && j.Code[i+5] == 'n' && j.Code[i+6] == ' ' {
			toks = append(toks, NewToken(RETURN, "return"))
			i += 6
		} else if i+5 < len(j.Code) && ac == 'f' && j.Code[i+1] == 'n' && j.Code[i+2] == ' ' {
			toks = append(toks, NewToken(FUNCTION, "function"))
			i += 2
		} else if i+6 < len(j.Code) && ac == 's' && j.Code[i+1] == 't' && j.Code[i+2] == 'r' && j.Code[i+3] == 'u' && j.Code[i+4] == 'c' && j.Code[i+5] == 't' && j.Code[i+6] == ' ' {
			toks = append(toks, NewToken(STRUCT, "struct"))
			i += 7
		} else if i+8 < len(j.Code) && ac == '@' && j.Code[i+1] == 'i' && j.Code[i+2] == 'm' && j.Code[i+3] == 'p' && j.Code[i+4] == 'o' && j.Code[i+5] == 'r' && j.Code[i+6] == 't' && j.Code[i+7] == ' ' {
			toks = append(toks, NewToken(IMPORT, "@import"))
			i += 8
		} else if i+3 < len(j.Code) && ac == 'v' && j.Code[i+1] == 'a' && j.Code[i+2] == 'r' && j.Code[i+3] == ' ' {
			toks = append(toks, NewToken(VAR, "var"))
			i += 3
		} else if i+4 < len(j.Code) && ac == '@' && j.Code[i+1] == 'p' && j.Code[i+2] == 'k' && j.Code[i+3] == 'g' && j.Code[i+4] == ' ' {
			toks = append(toks, NewToken(PACKAGE, "@pkg"))
			i += 5
		} else if i+3 < len(j.Code) && ac == '@' && j.Code[i+1] == 'c' && j.Code[i+2] == 'p' && j.Code[i+3] == 'p' {
			toks = append(toks, NewToken(CPP, "@cpp"))
			i += 4
		} else if i+3 < len(j.Code) && ac == 'p' && j.Code[i+1] == 'u' && j.Code[i+2] == 'b' && j.Code[i+3] == ' ' {
			toks = append(toks, NewToken(PUBLIC, "pub"))
			i += 3
		} else if i+4 < len(j.Code) && ac == 'p' && j.Code[i+1] == 'r' && j.Code[i+2] == 'i' && j.Code[i+3] == 'v' && j.Code[i+4] == ' ' {
			toks = append(toks, NewToken(PRIVATE, "priv"))
			i += 4
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
		} else if ac == ',' {
			toks = append(toks, NewToken(COMMA, "."))
			i++
		} else if ac == '%' {
			toks = append(toks, NewToken(PORCENT, "%"))
			i++
		} else if ac == '.' {
			toks = append(toks, NewToken(POINT, "."))
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
		} else if ac == '{' {
			toks = append(toks, NewToken(OPEN_BRACKET, "{"))

			i++
		} else if ac == '}' {
			toks = append(toks, NewToken(CLOSE_BRACKET, "}"))
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
