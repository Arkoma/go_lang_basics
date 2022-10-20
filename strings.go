// Strings
package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	// Strings are immutable in go so you can't do this
	// book[0] = 116

	// Slice (start, end) 0 based, half empty range like this
	fmt.Println(book[4:11])

	//Slice with no end ... from 4 to the end
	fmt.Println(book[4:])

	//Slice with no start ... up to 4
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was ½ price")

	// Multi line (it also ignores back slashes)
	poem := `
		The road goes ever on
		Down from the door where it began
		...
		`
	fmt.Println(poem)
}
