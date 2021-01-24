package main

import (
	"fmt"
	"math"
	"unicode/utf8"
	"unsafe"
)

const (
	Sunday = iota
	Monday
	Tueday
	_
	Sednesday
)

func main() {
	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxUint16

	fmt.Println(num1)
	fmt.Println(num2)

	fmt.Println("type size : ", unsafe.Sizeof(num1))
	fmt.Println("type size : ", unsafe.Sizeof(num2))

	var s1 string = "abcd"
	fmt.Println("str len : ", len(s1))
	fmt.Println("str len2 : ", utf8.RuneCountInString(s1))
	fmt.Printf("s1[0]: %c \n", s1[0])

	fmt.Println(Sunday, Monday, Tueday, Sednesday)

}
