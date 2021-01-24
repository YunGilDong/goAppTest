package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello(n int) {
	fmt.Println("Hello", n)
}

func hello1() {
	fmt.Println("Hello")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	once := new(sync.Once)

	//var fh func(n int) = hello

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine : ", i)
			once.Do(func() {
				hello(i)
			})
		}(i)
	}

	fmt.Scanln()
}
