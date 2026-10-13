package hello_world

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func HelloWorld() {
	fmt.Println("Hello from Go")
	fmt.Println("Bye")
}

func Run() {
	fmt.Println(Hello("World", ""))
}
