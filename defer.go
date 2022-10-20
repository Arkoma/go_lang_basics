// Errors
package main

import (
	"fmt"
)

func main() {
	worker()
}

func worker() {
	resource1, err := acquire("A")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer release(resource1)

	resource2, err := acquire("B")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer release(resource2)
	fmt.Println("worker")
} // you'll see defers are called in reverse order

func acquire(name string) (string, error) {
	return name, nil
}

func release(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}
