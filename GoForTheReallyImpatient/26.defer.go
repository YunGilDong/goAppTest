package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func HelloWorld() {
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("Hello ")
	}()
}

func main() {

	HelloWorld()

	defer world()
	hello()
	hello()
	hello()

}
