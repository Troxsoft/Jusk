package main

import (
	"encoding/json"
	"fmt"
	"jusklang/pkg"
	"os"
	"os/exec"
	"strings"
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
		k := strings.Split(os.Args[1][:len(os.Args[1])-2], "/")
		cmd := exec.Command("g++", os.Args[1][:len(os.Args[1])-2]+"cpp", "-o", k[len(k)-1]+"exe")
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
		//fmt.Println(jusk.Compile())
	}

}
