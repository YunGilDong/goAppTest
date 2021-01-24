package main

import (
	"fmt"
)

type ObjectA struct {
	val string
}

func main() {
	//var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a [5]int
	fmt.Println(a)

	c := [...]int{11, 22, 33, 44, 55, 66, 77}
	fmt.Println(c)

	for i, value := range c {
		fmt.Println(i, value)
	}

	arrObj := [5]ObjectA{ObjectA{val: "A"}, ObjectA{val: "B"}, ObjectA{val: "C"}, ObjectA{val: "D"}, ObjectA{val: "E"}}

	arrObj2 := arrObj

	fmt.Println(arrObj, arrObj2)
	fmt.Printf("%p %p", &arrObj, &arrObj2)

}
