package main

import "fmt"

const taxRate float32 = 5.0

func main() {
	calculatorMain()
}

func calculatorMain() {
	fmt.Println("Profit Calculator")

	var revenue float32
	var expenses float32

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	ebt, profit := calculator(revenue, expenses)

	fmt.Println("ebt, profit", ebt, profit)
}

func calculator(revenue float32, expenses float32) (float32, float32) {
	ebt := revenue - expenses
	//fmt.Println("EBT", ebt)

	profit := ebt * (1 - (taxRate / 100))
	//fmt.Println("Profit", profit)

	ratio := (ebt / profit)
	Final := fmt.Sprintln("Ratio", ratio)

	fmt.Print(Final)

	return ebt, profit
}
