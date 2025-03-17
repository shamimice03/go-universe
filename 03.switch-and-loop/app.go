package main

import "fmt"

func main() {
	accountInfoSwitch()
}

func accountInfoSwitch() {

	balance := 2000

	for i := 0; i <= 10; i++ {
		fmt.Println("\n--- Your Bank Account ---")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Withdraw Funds")
		fmt.Println("*. Exit")
		fmt.Print("\nPlease select an option (1-2) or press any key to exit: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", balance)
		case 2:
			var withdrawAmount int
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if balance >= withdrawAmount {
				balance -= withdrawAmount
				fmt.Println("Updated balance:", balance)
			} else {
				fmt.Println("Insufficient balance, current account balance:", balance)
				return
			}
		default:
			print("Exit!\nThank you, see you soon!\n")
			return
		}
	}
}
