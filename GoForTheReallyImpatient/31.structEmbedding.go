package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) greeting() {
	fmt.Println("Hello~~")
}

type Student struct {
	Person
	school string
	grade  int
}

func (s *Student) greeting() {
	fmt.Println("Hello Students~")
}

func main() {
	var s *Student = new(Student)
	s.Person.greeting()

	var s2 Student
	s2.greeting()

}
