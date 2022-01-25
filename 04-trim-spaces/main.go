package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl04 := `Hi {{- . -}}  , how are you?
`
	t, err := template.New("tpl04").Parse(tpl04)
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	err = t.Execute(os.Stdout, "Johnny")
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}
}
