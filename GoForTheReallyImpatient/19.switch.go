package main

import (
	"fmt"
)

func main() {
	i := 2
	switch {
	case i%2 == 1:
		fmt.Println("odd")
	case i%2 == 0:
		fmt.Println("not odd")
	default:
		fmt.Println("def")
	}
}
