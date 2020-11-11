package fizzbuzz

func Sum(number int) string {
	var answer string

	if number%3 == 0 {
		answer = "Fizz"
	} else if number%5 == 0 {
		answer = "Buzz"
	}

	return answer
}
