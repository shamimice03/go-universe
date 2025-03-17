package main

import "fmt"

type Product struct {
	title string
	id    string
	price int
}

func main() {

	prices := [6]float64{1, 2, 3, 4, 5, 6}

	fmt.Println("prices", prices)

	subsetPrices := prices[:2]
	fmt.Println("subnetPrices", subsetPrices)

	// we can still take prefix elements
	test := subsetPrices[:3]
	fmt.Println("test", test)

	fmt.Println("prices: ", len(prices), cap(prices))
	fmt.Println("subnetPrices: ", len(subsetPrices), cap(subsetPrices))
	fmt.Println("test: ", len(test), cap(test))

	//dynamic array
	var elements []int

	for i := 0; i <= 10; i++ {
		elements = append(elements, i)
	}

	fmt.Println(elements)

	// delete from front
	elements = elements[1:]
	fmt.Println(elements)

	elements = elements[:len(elements)-2]
	fmt.Println(elements)

	// array of struct
	var products []Product

	product1 := Product{
		title: "rice",
		id:    "123",
		price: 500,
	}

	product2 := Product{
		title: "fish",
		id:    "234",
		price: 900,
	}

	products = append(products, product1, product2)

	fmt.Println(products)

	// unpacking ( merging two array )
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{100, 200, 300, 400, 500}

	arr1 = append(arr1, arr2...)
	fmt.Println(arr1)

	// make
	len := 2
	cap := 10
	arr := make([]string, len, cap)
	// for len empty value will be initilized
	arr = append(arr, "test") // this will be added on 3rd index, since first two are empty
	fmt.Println(arr)

	arr[0] = "hwllp"
	fmt.Println(arr)

	for index, value := range arr1 {
		fmt.Println(index, value)
	}

	for range arr1 {
		fmt.Println()
	}

}

// outputs
/*
prices [1 2 3 4 5 6]
subnetPrices [1 2]
test [1 2 3]
prices:  6 6
subnetPrices:  2 6
test:  3 6
[0 1 2 3 4 5 6 7 8 9 10]
[1 2 3 4 5 6 7 8 9 10]
[1 2 3 4 5 6 7 8]
[{rice 123 500} {fish 234 900}]
[1 2 3 4 5 100 200 300 400 500]
[  test]
[hwllp  test]
0 1
1 2
2 3
3 4
4 5
5 100
6 200
7 300
8 400
9 500
*/
