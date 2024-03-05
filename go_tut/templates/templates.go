package templates

import (
	"log"
	"os"
	"text/template"
)

func Run() {
	// t, err := template.New("test").Parse("Hello, {{.Name}}!")
	t, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	// data := struct {
	// 	Name string
	// }{
	// 	Name: "World",
	// }
	// if err := t.Execute(os.Stdout, nil); err != nil {
	// 	log.Fatal("Execute: ", err)
	// }
	slices := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	if err := t.ExecuteTemplate(os.Stdout, "template.gohtml", slices); err != nil {
		log.Fatal("Execute: ", err)
	}
}
