package main

import "fmt"

/*// part 1.
func main() {
	fmt.Printf("Hello, world")
}*/

/*// part 2.
func Hello() string {
	return "Hello, world"
}*/

/*// part 3.
func Hello(name string) string {
	return "Hello, " + name
}*/

/*//part 4.
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}*/

/*//part 5.
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"

//part 6. - adding french
const frenchHelloPrefix = "Bonjour, "
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}*/

//part 7.
/*const spanish = "Spanish"
const french = "French"
const sweden = "Sweden"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const swedishHelloPrefix = "Hallå, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case sweden:
		prefix = swedishHelloPrefix
	}
	return prefix + name
}*/

// part 8.
const (
	spanish = "Spanish"
	french = "French"
	sweden = "Sweden"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	swedishHelloPrefix = "Hallå, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case sweden:
		prefix = swedishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
/* Function signature includes a named return value (prefix string).
	This will create a variable called prefix in your function.
		- It will be assigned the "zero" value. This depends on the type, 
		for example ints are 0 and for strings it is "".
   The function name (greetingPrefix) starts with a lowercase letter. 
   In Go, public functions start with a capital letter and private ones start with a lowercase. 
   We don't want the internals of our algorithm to be exposed to the world, so we made this function private.
   Also, we can group constants in a block instead of declaring them each on their own line. 
   It's a good idea to use a line between sets of related constants for readability. 
*/

func main() {
	fmt.Println(Hello("world", ""))
}