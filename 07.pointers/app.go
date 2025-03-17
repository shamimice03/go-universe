package main

import "fmt"

func main() {
	age := 20

	new_age := &age

	fmt.Println("Value of age variable:", age)
	fmt.Println("Address of age variable:", &age)
	fmt.Println("Address of new_age variable", new_age)
	fmt.Println("Value of new age", *new_age) // dereferencing

	*new_age = 5

	fmt.Println("Value of age variable, after chaning new_age:", age)

	fmt.Println("Combained age:", combainedAge(&age))
}

func combainedAge(age *int) int {
	fmt.Println("address of age, inside function", age)
	*age = *age + 10
	return *age
}

// How pointer works
/*
func main(){
	value := 5
	address := &value
	combainedAge(address)
}

func combainedAge(address *int){
    value = *address // dereferencing
}
*/
