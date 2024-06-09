package main

/*
#cgo CFLAGS: -g -Wall
#include <stdio.h>
#include <stdlib.h>
void execCommand(const char * comand);
void execCommand(const char * comand){
	system(comand);
}
*/
import "C"
import (
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
	//fmt.Printf("%+v\n", jusk.Tokens)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		// b, _ := json.MarshalIndent(jusk.Tokens, "", "  ")
		// fmt.Printf("Tokens %+v\n", string(b))
		jusk.GenerateAst()

		// a, _ := json.MarshalIndent(jusk.Astes.Nodes, "", "   ")
		// fmt.Printf("Ast %+v\n", string(a))
		var pedro []string
		if strings.Contains(os.Args[1], "/") {

			pedro = strings.Split(os.Args[1], "/")
		} else {
			pedro = strings.Split(os.Args[1], "\\")

		}
		err = os.WriteFile(os.Args[1][:len(os.Args[1])-2]+"cpp", []byte(jusk.Compile(os.Args[1][:len(os.Args[1])-len(pedro[len(pedro)-1])])), 0777)
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("g++", os.Args[1][:len(os.Args[1])-2]+"cpp", "-o", os.Args[1][:len(os.Args[1])-3])
		//err = cmd.Run()

		r, err := cmd.CombinedOutput()
		fmt.Print(string(r))
		if err != nil {
			panic(cmd.String() + "   " + err.Error())
		}
		C.execCommand(C.CString(fmt.Sprintf("%s", os.Args[1][:len(os.Args[1])-3])))

		//fmt.Println(jusk.Compile())
	}

}
