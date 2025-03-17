package main

import (
	"fmt"
)

// func printAnything(value interface{}) {
// 	fmt.Println(value)
// }

func printAnything(value any) {
	fmt.Println(value)
}

func printSomething(value interface{}) {

	typedVal, ok := value.(int)

	if ok {
		fmt.Println(typedVal)
		fmt.Println("Integer")
	}

	/*
	   switch value.(type) {
	   case int:

	   	fmt.Println("Integer:", value)

	   case float64:

	   	fmt.Println("Float;", value)

	   default:

	   		fmt.Println("anything else")
	   	}
	*/
}

func main() {
	// printAnything(1)
	// printAnything(1.1)
	// printAnything("test")

	// printSomething(20)
	// printSomething(2.1)
	// printSomething("test")

}
