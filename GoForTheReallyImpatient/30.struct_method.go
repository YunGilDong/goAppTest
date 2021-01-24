package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) area() int {
	return r.width * r.height
}

func (r *Rectangle) scaleA(factor int) {
	r.width = r.width * factor
	r.height = r.height * factor
}

func (r Rectangle) scaleB(factor int) {
	r.width = r.width * factor
	r.height = r.height * factor
}

func main() {
	rect := &Rectangle{10, 20}
	fmt.Println(rect.area())

	rect1 := &Rectangle{30, 30}
	rect1.scaleA(10)
	fmt.Println(rect1)

	rect2 := &Rectangle{30, 30}
	rect2.scaleB(10)
	fmt.Println(rect2)

}
