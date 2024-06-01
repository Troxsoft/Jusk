package main

import (
	"encoding/json"
	"fmt"
	"jusklang/pkg"
	"os"
	"os/exec"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	f := os.Args[1]
	b, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	code := string(b)
	jusk := pkg.NewJuskLang(code)
	err = jusk.Tokenize()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		b, _ := json.MarshalIndent(jusk.Tokens, "", "  ")
		fmt.Printf("Tokens %+v\n", string(b))
		jusk.GenerateAst()

		a, _ := json.MarshalIndent(jusk.Astes.Nodes, "", "   ")
		fmt.Printf("Ast %+v\n", string(a))
		err = os.WriteFile(os.Args[1][:len(os.Args[1])-2]+"cpp", []byte(jusk.Compile()), 0777)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("g++", os.Args[1][:len(os.Args[1])-2]+"cpp", "-o", os.Args[1][:len(os.Args[1])-3])
		err = cmd.Run()

		if err != nil {
			panic(cmd.String() + "   " + err.Error())
		}
		//fmt.Println(jusk.Compile())
	}

}
