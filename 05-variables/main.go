package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl05 := `
Variables: así se crea una variable en un template
{{- $mynumber := 14 }}
El valor de la variable es: {{ $mynumber }}
`
	t, err := template.New("tpl04").Parse(tpl05)
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}
}
