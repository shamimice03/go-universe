package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var balance float64 = 2000

const balanceFile string = "balance.txt"

// How ReadFile works here:
// - ReadFile reads a file, and convert it to slice of byte[] ( data var here)
// - Then those slice byte[] can be converted to string
// - Then convert that string to float64
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile) // get slice of byte[]

	if err != nil {
		return 0, errors.New("Failed to find balance in file.")
	}

	balanceText := string(data)                         // convert slice of byte[] to string
	balance, err := strconv.ParseFloat(balanceText, 64) // convert string to float

	if err != nil {
		return 0, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)              // convert float64 to string
	balanceInBytes := []byte(balanceText)           // convert string to byte[] slice
	os.WriteFile(balanceFile, balanceInBytes, 0644) // write to file
}

func main() {
	accountInfo()
}

func accountInfo() {
	balance, err := getBalanceFromFile()

	if err != nil {
		panic(err)
	}

	for i := 0; i <= 10; i++ {

		fmt.Println("\n--- Your Bank Account ---")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Withdraw Funds")
		fmt.Println("3. Deposit Funds")
		fmt.Println("*. Exit")
		fmt.Print("\nPlease select an option (1-3) or press any key to exit: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your Balance: %.2f\n", balance)
		} else if choice == 2 {
			var withdrawAmount float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if balance >= withdrawAmount {
				balance -= withdrawAmount
				fmt.Printf("Your new balance: %.2f\n", balance)
				writeBalanceToFile(balance)
			} else {
				fmt.Println("Insufficient balance, current account balance:", balance)
				break
			}
		} else if choice == 3 {
			var depositAmount float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount >= 0 {
				balance += depositAmount
				fmt.Printf("Your new balance: %.2f\n", balance)
				writeBalanceToFile(balance)
			} else {
				println("Invalid ammount, please try again!")
				break
			}

		} else {
			print("Exit!\nThank you, see you soon!\n")
			return
		}
	}

}
