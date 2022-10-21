package main

import (
	"math"
	"testing"
)

func TestSquare(t *testing.T) {
	t.Run("test square with length of 7", func(t *testing.T) {
		s := Square{7}
		got := s.Area()
		want := 49.0
		assertFloatsEqual(t, got, want)
	})

	t.Run("test square with length of 10", func(t *testing.T) {
		s := Square{10}
		got := s.Area()
		want := 100.0
		assertFloatsEqual(t, got, want)
	})
}

func assertFloatsEqual(t testing.TB, got, want float64) {
	t.Helper() // needed so that test will report line number from function call instead of line number in the helper method
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestCircleAreaWithRadiusOf2(t *testing.T) {
	c := Circle{2}
	got := c.Area()
	want := math.Pi * 2 * 2
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
