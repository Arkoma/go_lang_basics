// Using params
package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4}
	/*
		This works because a slice parameter works like a
		pointer to one address in memory
	*/
	doubleValuesAt(values, 2)
	fmt.Println(values)

	/*
		This one doesn't because the int parameter is a
		new int entity ... not a pointer
	*/
	val := 10
	double(val) // returns 10 instead of expected 20
	fmt.Println(val)

	// how you can pass a pointer
	doublePtr(&val)
	fmt.Println(val)
}

func doubleValuesAt(values []int, index int) {
	values[index] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}
