package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	b := a[0:6:8]

	fmt.Println(a)
	fmt.Println(len(b), cap(b))
	fmt.Println(b)

}
