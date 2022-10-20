// Maps
package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.05,
		"MSFT": 287.70, // must use trailing comma
	}

	//Number of items returns 3
	fmt.Println(len(stocks))

	//Get a value
	fmt.Println(stocks["MSFT"])

	/*
		If you look up a key that isn't there you get
		the value type's zero value ... for example the
		float64 type's zero value is 0 so we expect to
		get a 0 here.
	*/
	fmt.Println(stocks["TSLA"])

	// Use two values to see if found
	value, there := stocks["TSLA"]
	if there {
		fmt.Println(value)
	} else {
		fmt.Println("TSLA not found")
	}

	// Set a value in a map
	stocks["TSLA"] = 822.12
	fmt.Println(stocks)

	// Delete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value for loop gives the keys of maps
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value for loop gives us keys AND values of maps
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
