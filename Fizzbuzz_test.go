package fizzbuzz

import "testing"

func TestSum(t *testing.T) {

	t.Run("Test for Fizz", func(t *testing.T) {

		got := Sum(3)
		want := "Fizz"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Test for Buzz", func(t *testing.T) {

		got := Sum(5)
		want := "Buzz"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
