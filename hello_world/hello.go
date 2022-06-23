package main

import "fmt"

const engHelloPrefix = "Hello, "
const spaHelloPrefix = "Hola, "
const freHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return prefixGreeting(lang) + name
}

func prefixGreeting(lang string) (prefix string) {
	switch lang {
	case "French":
		prefix = freHelloPrefix
	case "Spanish":
		prefix = spaHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", "English"))
}
