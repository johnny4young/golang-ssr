package main

import (
	"os"
	"text/template"
)

func main() {
	myFunctions := template.FuncMap{
		"double": func(x int) int {
			return x * 2
		},
		"triple": func(x int) int {
			return x * 3
		},
	}
	tpl, err := template.New("tpl09").Funcs(myFunctions).Parse("{{double .}} {{triple .}}")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, 5)
	if err != nil {
		panic(err)
	}

}
