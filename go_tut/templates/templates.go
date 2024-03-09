package templates

import (
	"os"
	"text/template"
)

func Run() {
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

	tmpl, err = template.New("test").Parse(`Name: {{.Name}}, Age: {{.Age}}
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
}
