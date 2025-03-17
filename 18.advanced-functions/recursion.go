package main

func factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}
