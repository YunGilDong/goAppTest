package main

import (
	"fmt"
	rf "reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func main() {
	p := Person{}
	ptrP := &p
	name, ok := rf.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := rf.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))

	fmt.Println(p, ptrP)
}
