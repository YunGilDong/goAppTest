package main

import (
	"fmt"
)

type IncreaseAble interface {
	Increase()
}
type IncreaseAble2 interface {
	Increase2()
}

type Number struct {
	val int
}

type Number2 struct {
	val int
}

func (v *Number) Increase() {
	v.val += 1
}

type Chracter struct {
	val string
}

func (v *Chracter) Increase() {
	v.val += "*"
}

func Increase2(v IncreaseAble2) {
	switch t := v.(type) {
	case int:
		fmt.Println("inter", t)
	}

}

func main() {

	numVal := Number{val: 1}
	chVal := Chracter{val: "*"}

	numVal.Increase()
	chVal.Increase()

	fmt.Println("NUM : ", numVal)
	fmt.Println("CH : ", chVal)

	num1 := Number{val: 2}
	num2 := Number2{val: 3}

	Increase2(num1)
	Increase2(num2)

	fmt.Println(num1, num2)

}
