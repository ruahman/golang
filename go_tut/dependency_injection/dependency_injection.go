package dependency_injection

import (
	"fmt"
	"io"
	"net/http"
)

// dependency injection dooes not need a framework,
// it is just a pattern that allows you to inject dependencies into your code
// DI allows you contol over the dependencies of your code and be able to mock them.
// is also helps to to decouple your code and make it more testable.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// http.ResponseWriter also implements io.Writer
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
