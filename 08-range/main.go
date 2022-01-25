package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	cursos := []string{"golang", "python", "ruby"}
	tpl08 := `
		<ul>
			{{- range .}}
				<li>{{- . -}}</li>
			{{- else }}
				<li>404 not found</li>
			{{- end}}
		</ul>
	`
	t, err := template.New("tpl08").Parse(tpl08)
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	err = t.Execute(os.Stdout, cursos)
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}
}
