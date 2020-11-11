package fizzbuzz

import "strconv"

func Sum(number int) string {
	var answer string

	if number%3 == 0 {
		answer = "Fizz"
	} else if number%5 == 0 {
		answer = "Buzz"
	} else {
		answer = strconv.Itoa(number)
	}

	return answer
}
