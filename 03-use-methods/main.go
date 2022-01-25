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

func (p *person) Hi() string {
	return "hi, I'm " + p.Name
}

func main() {
	tpl03 := `Hola {{ .Name }}, you are {{ .Age }} years old!, {{ .Hi }}
`
	t, err := template.New("tpl03").Parse(tpl03)
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	p := person{
		Name: "Johnny",
		Age:  30,
	}

	err = t.Execute(os.Stdout, &p)
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}
}
