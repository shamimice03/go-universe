package main

import (
	"errors"
	"fmt"
	"os"
)

//const taxRate float32 = 5.0

func main() {
	calculatorMain()
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Println(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func storeResultToFile(result string) error {
	fileName := "result.txt"
	result_data := []byte(result)
	err := os.WriteFile(fileName, result_data, 0644)

	if err != nil {
		return errors.New("failed To save result")
	}

	return nil
}

func calculatorMain() {
	fmt.Println("Profit Calculator")
	var revenue float64 = getUserInput("Revenue: ")
	var expenses float64 = getUserInput("Expenses: ")
	var taxRate float64 = getUserInput("taxRate: ")

	if revenue <= 0 {
		panic("Invalid revenue value")
	}
	if expenses <= 0 {
		panic("Invalid expenses value")
	}
	if taxRate <= 0 {
		panic("Invalid taxRate value")
	}

	calculator(revenue, expenses, taxRate)
}

func calculator(revenue float64, expenses float64, taxRate float64) (float64, float64) {
	ebt := revenue - expenses
	// fmt.Println("EBT", ebt)

	profit := ebt * (1 - (taxRate / 100))
	// fmt.Println("Profit", profit)

	ratio := (ebt / profit)

	result := fmt.Sprintf("EBT %.2f\nProfit %.2f\nRatio %.2f\n", ebt, profit, ratio)
	println(result)

	err := storeResultToFile(result)
	if err != nil {
		panic(err)
	}

	return ebt, profit
}
