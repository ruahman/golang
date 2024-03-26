package hello_world

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// this is the domain of our app
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

func Run() {
	// this is the sideffect of our app, the side effect is printing to the console
	// it is not the domain of our app
	// it is a good idea to keep the domain and side effects separate
	fmt.Println(Hello("World", ""))
}
