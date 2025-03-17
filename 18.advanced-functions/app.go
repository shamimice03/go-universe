package main

import (
	"fmt"
)

// function as parameter
// function as return
// anonymous function

type transformFn func(int) int

func main() {

	var values1 []int
	values2 := []int{1, 2, 3, 4}
	values1 = append(values1, values2...)

	doubled := transformNumbers(&values2, double) // passing func as parameter
	tripled := transformNumbers(&values1, triple) // passing func as parameter
	fmt.Println(values1, values2)
	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformer(&doubled) // getting func as return
	transformerFn2 := getTransformer(&tripled) // getting func as return

	doubledAgain := transformNumbers(&doubled, transformerFn1)
	tripledAgain := transformNumbers(&tripled, transformerFn2)

	fmt.Println(doubledAgain)
	fmt.Println(tripledAgain)

	// anonymous func
	doubled = transformNumbers(&values2, func(v int) int {
		return v * 2
	})

	fmt.Println(doubled)

	// closures
	quadruple := createTransformer(4)
	quintuple := createTransformer(5)

	fmt.Println(transformNumbers(&values1, quadruple)) // values, func
	fmt.Println(transformNumbers(&values1, quintuple)) // values, func

	// recursion
	fmt.Println(factorial(5))

	// variadic func
	fmt.Println(sumpUp("summing up", 1, 2, 66, 63))

	// arry as input
	fmt.Println(sumpUp("summing up", values1...))
}

// variadic func
func sumpUp(startingValue string, numbers ...int) int {
	fmt.Println(startingValue)
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// closures
func createTransformer(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}

func getTransformer(numbers *[]int) func(int) int {

	if (*numbers)[0] == 2 {
		return double
	} else {
		return triple
	}
}

// can be written as
// func getTransformer(numbers *[]int) transformFn {

// 	if (*numbers)[0] == 2 {
// 		return double
// 	} else {
// 		return triple
// 	}
// }

func transformNumbers(numbers *[]int, tranform func(int) int) []int {

	var transform []int
	for _, v := range *numbers {
		transform = append(transform, tranform(v))
	}

	return transform
}

func double(v int) int {
	return v * 2
}

func triple(v int) int {
	return v * 3
}
