package main

import (
	"fmt"

	"cloudterms.net/package-usage/fileops" // Internal Package
	"github.com/Pallinder/go-randomdata"   // Third-party Package
)

const balanceFile string = "balance.txt"

func display() {
	fmt.Println("\n--- Your Bank Account ---")
	fmt.Println("Contact us:", randomdata.Address())
	fmt.Println("Phone:", randomdata.PhoneNumber())
	fmt.Println("1. Check Balance")
	fmt.Println("2. Withdraw Funds")
	fmt.Println("3. Deposit Funds")
	fmt.Println("*. Exit")
	fmt.Print("\nPlease select an option (1-3) or press any key to exit: \n")

}

func accountInfo() {
	balance, err := fileops.GetBalanceFromFile(balanceFile)

	if err != nil {
		panic(err)
	}

	for i := 0; i <= 10; i++ {

		display()

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
				fileops.WriteBalanceToFile(balanceFile, balance)
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
				fileops.WriteBalanceToFile(balanceFile, balance)
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

func main() {
	accountInfo()
}
