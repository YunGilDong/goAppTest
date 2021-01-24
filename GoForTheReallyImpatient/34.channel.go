package main

import (
	"fmt"
)

func num(a, b int) <-chan int {
	out := make(chan int)
	fmt.Println("#1")
	go func() {
		fmt.Println("#2")
		// if <-out != 0 {
		// 	out <- a
		// 	out <- b
		// }

		fmt.Println("#2-1", out)
		close(out) // close한채널은 더이상 데이터를 넣을 수 없다. but 채널에 데이터가 있을경우 읽기는 가능하다.
		if <-out != 0 {
			out <- b
		}

	}()

	fmt.Println("#3")
	return (out)
}

func sum(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		for i := range c {
			r = r + i
		}
		out <- r
	}()
	return out
}

func closureTest() int {

	a := 10
	fmt.Println("@1")
	func() {
		fmt.Println("@2")
		a = a + 100
	}()

	fmt.Println("@3")

	return a
}

func main() {
	c := num(1, 2)
	fmt.Println("#4", c)
	out := sum(c)

	fmt.Println(<-out)

	a := closureTest()
	fmt.Println("@4")
	fmt.Println(a)
}
