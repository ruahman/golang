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
		Test   bool
		Name   string
		Num1   int
		Num2   int
	}
	var data Data = Data{slices, maps, true, "Andy", 10, 20}

	if err := t.ExecuteTemplate(os.Stdout, "template.gohtml", data); err != nil {
		log.Fatal("Execute: ", err)
	}
}
