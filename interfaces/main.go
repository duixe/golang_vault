package main

import "fmt"

//interface is used to define a set of methods and their return types such that any other types (struct) that has any of the defined methods
//in the interface can be called a type of that interface
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb);
	printGreeting(sb);
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting());
}

func (eb englishBot) getGreeting() string {
	return "Hi there !"
}

func (sb spanishBot) getGreeting() string {
	return "Hola !"
}