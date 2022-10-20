// Errors
package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

// all fields that start with a capital are exported (like public in Java)
// any fields that start with a lower case are not exported (or private)

func main() {
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)
	fmt.Printf("%#v\n", b1)
	fmt.Println(b1.CampaignID)

	b2 := Budget{
		Balance:    19.3,
		CampaignID: "puppies",
	}
	fmt.Printf("%#v\n", b2) // Expiration returns its zero value

	var b3 Budget
	fmt.Printf("%#v\n", b3) // all values return their zero value

}
