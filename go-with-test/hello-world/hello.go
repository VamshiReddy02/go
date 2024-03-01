package main

import "fmt"

const sayHello = "hello, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	if lang == "Spanish" {
		return "Hola, " + name
	}
	return sayHello + name
}

func main() {
	fmt.Println(Hello("Vamshi", "Spanish"))
}
