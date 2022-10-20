// Slices
package main

import (
	"fmt"
)

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// Length
	fmt.Println(len(loons))

	fmt.Println("----")
	// 0 indexing
	fmt.Println(loons[1])

	fmt.Println("----")
	// slices
	fmt.Println(loons[1:])

	fmt.Println("----")
	// for loop
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("----")
	// for loop using range
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("----")
	// for loop getting value and index
	for i, value := range loons {
		fmt.Printf("%s at %d\n", value, i)
	}

	fmt.Println("----")
	// for loop getting just the value
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("----")
	// append
	loons = append(loons, "elmer")
	fmt.Println(loons)
}
