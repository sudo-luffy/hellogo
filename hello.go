package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func greetngPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func Hello(name, language string) string {
	prefix := greetngPrefix(language)
	if name == "" {
		name = "world"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
