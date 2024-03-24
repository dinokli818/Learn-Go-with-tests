package main

import (
	"fmt"
)

const (
	chinese            = "Chinese"
	french             = "French"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	chineseHelloPrefix = "你好, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func main() {
	fmt.Println(Hello("world", ""))
}
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (Prefix string) {
	switch language {
	case chinese:
		Prefix = chineseHelloPrefix
	case french:
		Prefix = frenchHelloPrefix
	case spanish:
		Prefix = spanishHelloPrefix
	default:
		Prefix = englishHelloPrefix

	}
	return
}
