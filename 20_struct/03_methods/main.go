package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func (p person) SayHello() string {
	return "hello"
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

	var p person
	fmt.Println(p.SayHello())
}
