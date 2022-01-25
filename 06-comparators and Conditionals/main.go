package main

import (
	"log"
	"os"
	"text/template"
)

type numbers struct {
	X int
	Y int
}

func main() {
	tpl06 := `
Comparadores:
eq: Es igual a {{ .X }} es igual a {{ .Y }}? {{ if eq .X .Y }}SI{{ else }}NO{{ end }}
ne: Es diferente a {{ .X }} es diferente a {{ .Y }}? {{ if ne .X .Y }}SI{{ else }}NO{{ end }}
lt: Es menor que {{ .X }} es menor que {{ .Y }}? {{ if lt .X .Y }}SI{{ else }}NO{{ end }}
gt: Es mayor que {{ .X }} es mayor que {{ .Y }}? {{ if gt .X .Y }}SI{{ else }}NO{{ end }}
le: Es menor o igual que {{ .X }} es menor o igual que {{ .Y }}? {{ if le .X .Y }}SI{{ else }}NO{{ end }}
ge: Es mayor o igual que {{ .X }} es mayor o igual que {{ .Y }}? {{ if ge .X .Y }}SI{{ else }}NO{{ end }}
`

	t, err := template.New("tpl06").Parse(tpl06)
	if err != nil {
		log.Fatalf("Error parsing template	%v", err)
	}

	n := numbers{5, 6}
	err = t.Execute(os.Stdout, n)
	if err != nil {
		log.Fatalf("Error executing template	%v", err)
	}
}
