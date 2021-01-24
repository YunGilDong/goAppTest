package main

import "fmt"

type Obj struct {
	val string
}

func hello(n *int) {
	*n = 2
}

func InitObj() *Obj {
	return &Obj{val: "OBJ"}
}

func main() {
	pObj := InitObj()
	fmt.Printf("%p\n", pObj)
	fmt.Println(*pObj)

}
