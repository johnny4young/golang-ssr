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

	t, err := template.New("").ParseGlob("templates/*.tpl")
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	p := person{
		Name: "Gopher",
		Age:  19,
	}
	err = t.ExecuteTemplate(os.Stdout, "invitation.tpl", p)
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}
}
