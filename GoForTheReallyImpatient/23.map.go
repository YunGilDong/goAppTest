package main

import (
	"fmt"
)

type ObjectB struct {
	val string
}

func main() {
	a := make(map[string]int)
	a["Hello"] = 10
	a["world"] = 20

	fmt.Println(a["Hello"])
	fmt.Println(a["world"])
	fmt.Println(a["aa"])

	b := make(map[string]ObjectB)

	b["A"] = ObjectB{val: "AA"}
	b["B"] = ObjectB{val: "BB"}

	fmt.Println(b["A"])

	if value, ok := b["AB"]; ok {
		fmt.Println(b["AB"], value)
	}

	delete(b, "A")
	for k, v := range b {
		fmt.Println(k, v)
	}

}
