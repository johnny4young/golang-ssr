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

	t, err := template.New("invitation.tpl").ParseFiles("templates/invitation.tpl")
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	p := person{
		Name: "Gopher",
		Age:  19,
	}
	err = t.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}
}
