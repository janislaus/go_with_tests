package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	german             = "German"
	english            = "English"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	germanHelloPrefix  = "Hallihallo, "
)

func Hello(name string, lang string) string {

	prefixes := map[string]string{
		english: englishHelloPrefix,
		spanish: spanishHelloPrefix,
		french:  frenchHelloPrefix,
		german:  germanHelloPrefix}

	if name == "" {
		return prefixes[lang] + "World!"
	}
	return prefixes[lang] + name
}

func main() {
	fmt.Println(Hello("", "English"))
}
