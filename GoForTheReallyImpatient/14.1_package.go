package main

import (
	"fmt"
	f "fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU Count : ", runtime.NumCPU())
	f.Println("CPU Count : ", runtime.NumCPU())

}
