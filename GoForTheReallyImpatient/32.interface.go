package main

import "fmt"

type MyInt int

func (i MyInt) Print() {
	fmt.Println(i)
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print()
}

func DataPrint(p Printer) {
	p.Print()
}

func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}

	pArr := []Printer{i, r}
	for index, _ := range pArr {
		pArr[index].Print()
	}

	for _, value := range pArr {
		value.Print()
	}

}
