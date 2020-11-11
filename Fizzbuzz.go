package fizzbuzz

func Sum(number int) string {
	var answer string

	if number%3 == 0 {
		answer = "Fizz"
	}

	return answer
}
