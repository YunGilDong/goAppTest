package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

func (s *StructA) String() string {
	return "Val : " + s.val
}

type StructB struct {
	val int
}

func (s *StructB) String() string {
	return "StrucdtB : " + strconv.Itoa(s.val)
}

type Printalbe interface {
	String() string
}

func Println(p Printalbe) {
	fmt.Println(p.String())
}

func main() {
	a := &StructA{val: "AAA"}
	fmt.Println(a)

	b := &StructB{val: 10}
	fmt.Println(b)

}
