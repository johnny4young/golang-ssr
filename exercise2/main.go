package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func main() {
	tpl02 := `Hola {{ .Name }}, you are {{ .Age }} years old!
`
	t, err := template.New("tpl02").Parse(tpl02)
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	p := person{
		Name: "Johnny",
		Age:  30,
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}
}
