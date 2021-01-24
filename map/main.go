package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)

	m["bcd"] = "ccc"
	m["abc"] = "bbb"

	fmt.Println("bcd : ", m["bcd"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println(m1[53])

	fmt.Println("m1[55]=>", m1[55])

	m2 := make(map[int]int)
	m2[4] = 4
	fmt.Println("m2[10] : ", m2[10])

	m2[5] = 0
	fmt.Println("m2[5]>", m2[5])

	v := m2[5]
	v1 := m2[4]
	v2, ok := m2[10]
	fmt.Println(v, v1, v2, ok)

	delete(m2, 5)

	v, ok1 := m2[5]
	fmt.Println(v, v1, v2, ok1)

	m2[2] = 2
	m2[10] = 10

	for key, value := range m2 {
		fmt.Println(key, " = ", value)
	}

}
