package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(newName string) {
	t.name = newName
}

func (t *Student) SetAge(age int) {
	t.age = age
}

func SetName(t *Student, name string) {
	t.name = name
}

func PrintStudent(u *Student) {
	fmt.Println(u)
}

func main() {
	var a *Student
	a = &Student{"aaa", 20, 10}

	SetName(a, "bbb")
	a.SetName("bbb")
	a.SetAge(30)
	PrintStudent(a)
}
