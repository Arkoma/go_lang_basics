package main

import (
	"testing"
)

func TestGenericsMinimum(t *testing.T) {
	t.Run("test returns error when sent an empty slice", func(t *testing.T) {
		_, err := minimum([]string{})
		assertError(t, err)
	})
	t.Run("test returns expected error message when sent an empty slice", func(t *testing.T) {
		_, err := minimum([]string{})
		assertErrorMessage(t, err, "cannot get minimum of empty slice")
	})
	t.Run("test returns no error from [3, 0, 4]", func(t *testing.T) {
		_, err := minimum([]int{3, 0, 4})
		assertNoError(t, err)
	})
	t.Run("test returns 1 from [3, 1, 4]", func(t *testing.T) {
		result, _ := minimum([]int{3, 1, 4})
		assertEquals(t, 1, result)
	})
	t.Run("test returns 1 from [3, 4, 1]", func(t *testing.T) {
		result, _ := minimum([]int{3, 4, 1})
		assertEquals(t, 1, result)
	})
	t.Run("test returns A from [\"B\", \"C\", \"A\"]", func(t *testing.T) {
		result, _ := minimum([]string{"B", "C", "A"})
		assertEquals(t, "A", result)
	})
	t.Run("test returns -2.3 from [567.4, -2.3, 0.0]", func(t *testing.T) {
		result, _ := minimum([]float64{567.4, -2.3, 0.0})
		assertEquals(t, -2.3, result)
	})

}

type Result interface {
	int | float64 | string
}

func assertEquals[T Result](t testing.TB, expected T, actual T) {
	t.Helper()
	if expected != actual {
		t.Error("actual result did not equal expected result")
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected error but got none")
	}
}

func assertErrorMessage(t testing.TB, err error, message string) {
	t.Helper()
	if err == nil {
		t.Error("expected error but got none")
	} else if err.Error() != message {
		t.Error("expected " + message + " but got " + err.Error())
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("did not expect error but got" + err.Error())
	}
}
