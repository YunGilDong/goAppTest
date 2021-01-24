package main

import (
	"fmt"
	io "io/ioutil"
)

func main() {

	if b, err := io.ReadFile("./hello.txt"); err == nil {
		fmt.Printf(">%s", b)
	}
}
