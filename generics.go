// Budget struct with a constructor
package main

import "errors"

type Ordered interface {
	int | float64 | string
}

func minimum[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, errors.New("cannot get minimum of empty slice")
	}
	minimum := values[0]
	for _, value := range values[1:] {
		if value < minimum {
			minimum = value
		}
	}
	return minimum, nil
}

func main() {

}
