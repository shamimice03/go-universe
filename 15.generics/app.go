package main

// this will not work
// func add(a, b any) any {
// 	return a + b
// }

func add[T int | string | float64](a, b T) T {
	result := a + b
	println(result)
	return result
}

func main() {

	add(2, 3)
	add(2.3, 3)
	add("md", "shamim")

}
