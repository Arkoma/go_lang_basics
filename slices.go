// Calculate maximal value in a slice
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 23, 15}
	maxnum := nums[0]

	for _, num := range nums[1:] {

		if num > maxnum {
			maxnum = num
		}

	}

	fmt.Println(maxnum)
}
