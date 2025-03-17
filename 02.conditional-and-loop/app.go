package main

import "fmt"

func main() {

	accountInfo()
}

func accountInfo() {
	balance := 2000

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
			fmt.Println("Your balance is:", balance)
		} else if choice == 2 {
			var withdrawAmount int
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if balance >= withdrawAmount {
				balance -= withdrawAmount
				fmt.Println("Updated balance:", balance)
			} else {
				fmt.Println("Insufficient balance, current account balance:", balance)
				break
			}
		} else if choice == 3 {
			var depositAmount int
			fmt.Print("Deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount >= 0 {
				balance += depositAmount
				println("Your new balance:", balance)
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
