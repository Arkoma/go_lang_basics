// Get content-type from sites
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//ch <- 353 // send ... channel will panic because it's not in a routine

	go func() {
		ch <- 353 // send
	}()

	val := <-ch //receive

	fmt.Printf("got %d\n", val)

	fmt.Println("----")

	const count = 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}() // loop is blocked until the reciever receives it

	for i := 0; i < count; i++ {
		val := <-ch // this is blocked until a value is sent on the receiver
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("----")

	// waits for close to signal we're done
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}

	fmt.Println("----")

	ch2 := make(chan int, 1) // buffered channel
	ch2 <- 19                // a buffered channel can recieve items equivalent to the buffer without being blocked
	val2 := <-ch2
	fmt.Println(val2)

}
