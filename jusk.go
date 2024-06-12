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
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func showInfo() {

	fmt.Printf(`		%s

	Version: %s
   ---
Building for:%s

`, color.HiMagentaString("Jusk Programing Language"), color.CyanString("%v", pkg.VERSION), color.YellowString("%s %s", runtime.GOOS, runtime.GOARCH))

}
func main() {
	// defer func() {
	// 	if err := recover(); err != nil {

	// 		fmt.Printf("Error: %s\n", color.RedString("%v", err))
	// 	}
	// }()
	app := &cli.App{
		Name:                 "Jusk",
		Description:          "The jusk programming language",
		Version:              pkg.VERSION,
		EnableBashCompletion: true,
		Usage:                "The Jusk Programing Language",
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "cpp",
				Aliases: []string{"c"},
				Usage:   "Translate jusk code to c++",

				Action: func(ctx *cli.Context) error {
					showInfo()
					nameFile := ctx.Args().First()
					fileB, err := os.ReadFile(nameFile)
					if err != nil {
						return err
					}
					jusk := pkg.NewJuskLang(string(fileB))
					err = jusk.Tokenize()
					if err != nil {
						return err
					}
					err = jusk.GenerateAst()
					if err != nil {
						return err
					}

					nameFile = strings.ReplaceAll(nameFile, "\\", "/")
					nameFile = strings.ReplaceAll(nameFile, "./", "")
					pedro := strings.Split(nameFile, "/")
					err = os.WriteFile(nameFile[:len(nameFile)-2]+"cpp", []byte(jusk.Compile(nameFile[:len(nameFile)-len(pedro[len(pedro)-1])], nil)), 0777)

					if err != nil {
						return err
					}

					fmt.Printf("Build(c++ code): %s %s\n", color.GreenString("%s", nameFile), color.YellowString("sucessfully"))
					return nil
				},
			},
			&cli.Command{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Build and run jusk file",

				Action: func(ctx *cli.Context) error {
					showInfo()
					nameFile := ctx.Args().First()
					fileB, err := os.ReadFile(nameFile)
					if err != nil {
						return err
					}
					jusk := pkg.NewJuskLang(string(fileB))
					err = jusk.Tokenize()
					if err != nil {
						return err
					}
					err = jusk.GenerateAst()
					if err != nil {
						return err
					}

					nameFile = strings.ReplaceAll(nameFile, "\\", "/")
					nameFile = strings.ReplaceAll(nameFile, "./", "")
					pedro := strings.Split(nameFile, "/")
					err = os.WriteFile(nameFile[:len(nameFile)-2]+"cpp", []byte(jusk.Compile(nameFile[:len(nameFile)-len(pedro[len(pedro)-1])], nil)), 0777)

					if err != nil {
						return err
					}
					cmd := exec.Command("g++", nameFile[:len(nameFile)-2]+"cpp", "-o", nameFile[:len(nameFile)-3])

					//err = cmd.Run()

					r, err := cmd.CombinedOutput()

					if err != nil {
						return fmt.Errorf("Build: %s with errors ! %s\n", color.GreenString("%s", nameFile), color.RedString(string(r)))
					}

					fmt.Printf("Build(executable): %s %s\n", color.GreenString("%s", nameFile), color.YellowString("sucessfully"))
					os.Remove(nameFile[:len(nameFile)-2] + "cpp")
					var execD string
					if runtime.GOOS == "windows" {
						execD = nameFile[:len(nameFile)-2] + "exe"
						execD = strings.ReplaceAll(execD, "/", "\\")
						C.execCommand(C.CString(fmt.Sprintf(".\\%s", execD)))
					} else {
						execD = nameFile[:len(nameFile)-2]
						C.execCommand(C.CString(fmt.Sprintf("./%s", execD)))
					}

					os.Remove(execD)
					return nil
				},
			},
			&cli.Command{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build a jusk file to executable",

				Action: func(ctx *cli.Context) error {
					showInfo()
					nameFile := ctx.Args().First()
					fileB, err := os.ReadFile(nameFile)
					if err != nil {
						return err
					}
					jusk := pkg.NewJuskLang(string(fileB))
					err = jusk.Tokenize()
					if err != nil {
						return err
					}
					err = jusk.GenerateAst()
					if err != nil {
						return err
					}

					nameFile = strings.ReplaceAll(nameFile, "\\", "/")
					nameFile = strings.ReplaceAll(nameFile, "./", "")
					pedro := strings.Split(nameFile, "/")
					err = os.WriteFile(nameFile[:len(nameFile)-2]+"cpp", []byte(jusk.Compile(nameFile[:len(nameFile)-len(pedro[len(pedro)-1])], nil)), 0777)

					if err != nil {
						return err
					}
					cmd := exec.Command("g++", nameFile[:len(nameFile)-2]+"cpp", "-o", nameFile[:len(nameFile)-3])

					//err = cmd.Run()

					r, err := cmd.CombinedOutput()

					if err != nil {
						return fmt.Errorf("Build: %s with errors ! %s\n", color.GreenString("%s", nameFile), color.RedString(string(r)))
					}

					fmt.Printf("Build(executable): %s %s\n", color.GreenString("%s", nameFile), color.YellowString("sucessfully"))
					os.Remove(nameFile[:len(nameFile)-2] + "cpp")
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
