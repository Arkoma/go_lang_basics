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
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("test square with length of 10", func(t *testing.T) {
		s := Square{10}
		got := s.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}

func TestCircleAreaWithRadiusOf2(t *testing.T) {
	c := Circle{2}
	got := c.Area()
	want := math.Pi * 2 * 2
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
