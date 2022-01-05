package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl01 := "Hello, world"
	t, err := template.New("tpl01").Parse(tpl01)
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}

}
