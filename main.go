package main

import "fmt"

type A struct {
	year int
}

func (a A) Greet() {
	fmt.Println("Greeting", a.year)
}

// embeding the type A in B, so B has the same greet method, and B can define a different Greet() method
type B struct {
	A
}

func main() {
	var a A
	a.year = 2021

	var b B
	b.year = 2022

	a.Greet()
	b.Greet()
}
