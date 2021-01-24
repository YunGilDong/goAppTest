package main

import "fmt"

func main() {
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}

	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println("Hellow~~")
}
