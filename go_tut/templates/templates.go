package templates

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Data3 struct {
	Name string
}

func (d Data3) Hello() string {
	return "Hello, " + d.Name
}

func Run() {
	// simple template
	tmpl, err := template.New("test").Parse("Hello, {{.}}!\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, "World")
	if err != nil {
		panic(err)
	}

	type Data struct {
		Name    string
		Age     int
		Sex     string
		Hobbies []string
	}

	data := Data{"John", 30, "Male", []string{"reading", "running", "swimming"}}

	tmpl, err = template.New("test2").Parse(`Name: {{.Name}}, Age: {{.Age}}
  {{ $name := .Name }}
  hobies: {{range .Hobbies}}
  {{$name}} likes {{.}}
  {{end}} 
  `)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

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

	type Data2 struct {
		Slices []string
		Maps   map[string]string
	}
	var data2 Data2 = Data2{slices, maps}

	if err := t.ExecuteTemplate(os.Stdout, "template.gohtml", data2); err != nil {
		log.Fatal("Execute: ", err)
	}

	// nested templates
	// tmpl, err = template.ParseGlob("*.gohtml")

	data3 := Data3{"Diego"}

	tmpl, err = template.ParseFiles("index.gohtml", "side.gohtml")
	if err != nil {
		panic(err)
	}
	if err = tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", data3); err != nil {
		log.Fatal("Execute: ", err)
	}
}
