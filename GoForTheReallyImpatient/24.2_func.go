package main

import "fmt"

func SumAndDiff(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff
}

func funcSum(a int, b int) int {
	return a + b
}

func main() {
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(sum, diff)

	var hello func(a int, b int) int = funcSum
	world := funcSum

	fmt.Println(hello(1, 2))
	fmt.Println(world(1, 2))

	func(s string) {
		fmt.Println("Hello world!", s)
	}("aaa")

}
