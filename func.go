// Functions
package main

import (
	"fmt"
)

func main() {
	val := add(1, 2)
	fmt.Println(val)

	times, remainder := divideAndGetModulo(7, 2)
	fmt.Printf("times=%d, remainder=%d", times, remainder)
}

func add(a int, b int) int {
	return a + b
}

func divideAndGetModulo(a int, b int) (int, int) {
	return a / b, a % b
}
