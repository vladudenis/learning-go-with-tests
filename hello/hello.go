package main

import "fmt"

const (
	French  string = "French"
	Spanish string = "Spanish"

	EnglishPrefix string = "Hello, "
	FrenchPrefix  string = "Bonjour, "
	SpanishPrefix string = "Hola, "
)

// Private function -> first letter is lowercase
func greetingPrefix(language string) (prefix string) {
	switch language {
	case French:
		prefix = FrenchPrefix
	case Spanish:
		prefix = SpanishPrefix
	default:
		prefix = EnglishPrefix
	}
	return
}

// Hello Public function -> first letter is uppercase
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
