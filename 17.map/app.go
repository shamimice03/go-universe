package main

import "fmt"

type stringIntMap map[string]int

func main() {
	// map[key_type]value_type
	// with initialization 0
	scores := make(map[string]int)

	// with initialization of 100 map
	userIDs := make(map[string]int, 100)

	// alternative way
	scores1 := map[string]int{
		"john":  10,
		"jason": 20,
	}

	scores1["helo"] = 20

	fmt.Println(scores)
	fmt.Println(userIDs)
	fmt.Println(scores1)

	// delete
	delete(scores1, "john")
	fmt.Println(scores1)

	test := make(stringIntMap)
	test["hello"] = 2
	fmt.Println(test)

	//loop

	for index, value := range scores1 {
		fmt.Println(index, value)
	}

}
