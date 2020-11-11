package fizzbuzz

import "testing"

func TestSum(t *testing.T) {

	t.Run("Test for Fizz", func(t *testing.T) {
		got := Sum(3)
		want := "Fizz"
		assertEquals(t, got, want)
	})

	t.Run("Test for Buzz", func(t *testing.T) {
		got := Sum(5)
		want := "Buzz"
		assertEquals(t, got, want)
	})

	t.Run("Test for normal integer", func(t *testing.T) {
		got := Sum(4)
		want := "4"
		assertEquals(t, got, want)
	})

	t.Run("Test for Fizzbuzz", func(t *testing.T) {
		got := Sum(15)
		want := "Fizzbuzz"
		assertEquals(t, got, want)
	})

}

func assertEquals(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
