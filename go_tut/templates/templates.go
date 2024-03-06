package templates

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func Run() {
	funcMap := template.FuncMap{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"cus": func(s string) string {
			return "cus: " + s
		},
	}
	t, err := template.New("template.gohtml").Funcs(funcMap).ParseFiles("template.gohtml")
	if err != nil {
		log.Fatal("Parse: ", err)
	}

	slices := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	maps := map[string]string{
		"Japan":  "Tokyo",
		"China":  "Beijing",
		"Canada": "Ottawa",
		"USA":    "Washington D.C.",
	}

	type Data struct {
		Slices []string
		Maps   map[string]string
	}
	var data Data = Data{slices, maps}

	if err := t.ExecuteTemplate(os.Stdout, "template.gohtml", data); err != nil {
		log.Fatal("Execute: ", err)
	}
}
