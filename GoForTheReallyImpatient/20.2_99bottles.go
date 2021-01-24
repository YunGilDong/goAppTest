package main

import (
	"fmt"
)

func main() {
	for bottles := 99; bottles > 0; bottles-- {
		switch {
		case bottles >= 2:
			fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", bottles, bottles)
			fmt.Print("Take one down, pass it around, ")
		case bottles == 1:
			fmt.Printf("%d bottle of beer on the wall, %d bottle of beer.\n", bottles, bottles)
			fmt.Print("Take one down, pass it around, No more bottles of beer on the wall\n")
			fmt.Print("Goto the stroe......")
		}

	}

}
