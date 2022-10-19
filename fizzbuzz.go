/*
Fizzbuzz challenge
Print the numbers between 1 to 20, however
- if the number is divisible by 3, print "fizz" instead
- if the number is divisible by 5, print "buzz" instead
- if the number is divisible by both 3 and 5, print "fizz buzz" instead
- Otherwise print the number
*/
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 21; i++ {
		// if i%3 == 0 && i%5 == 0 {
		// 	fmt.Println("fizz buzz")
		// 	continue
		// }

		// if i%3 == 0 {
		// 	fmt.Println("fizz")
		// 	continue
		// }

		// if i%5 == 0 {
		// 	fmt.Println("buzz")
		// 	continue
		// }

		// fmt.Println(i)

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizz buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
	// fmt.Println(1 % 5)
	// fmt.Println(7 % 5)
	// fmt.Println(10 % 5)
}
