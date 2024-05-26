package main

import (
	"encoding/json"
	"fmt"
	"jusklang/pkg"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	code := "var x:Str = \"Hola Mundo\""
	jusk := pkg.NewJuskLang(code)
	err := jusk.Tokenize()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		b, _ := json.MarshalIndent(jusk.Tokens, "", "  ")
		fmt.Printf("Tokens %+v\n", string(b))
		jusk.GenerateAst()

		a, _ := json.MarshalIndent(jusk.Astes.Nodes, "", "   ")
		fmt.Printf("Ast %+v\n", string(a))

		fmt.Println(jusk.Compile())
	}

}
