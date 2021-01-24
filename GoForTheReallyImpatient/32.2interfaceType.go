package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)
	default:
		return "error"
	}
}

func main() {
	fmt.Println(formatString(Person{"maria", 20}))
	fmt.Println(formatString(&Person{"maria", 20}))

	var andrew = new(Person)
	andrew.name = "Andrew"
	andrew.age = 35
	fmt.Println(formatString(andrew))
}
