package fizzbuzz

import "testing"

func TestSum(t *testing.T) {

	t.Run("Test for Fizz", func(t *testing.T) {

		got := Sum(3)
		want := "Fizz"

		if got != want {
			t.Errorf("got %d want %d given, %s", got, want)
		}
	})

}