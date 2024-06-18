package bankops

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


const balanceTxt = "balance.txt"
// Writing balance to file
func writeBalanceToFile(accountBalance float64){
	file, err := os.Create(balanceTxt)
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Fprintf(file, "Balance as of %s is $%.2f\n", formattedTime, accountBalance)
	}
}
// Reading balance from a file
func readBalanceFromFile() float64 {
	data, _ := os.ReadFile(balanceTxt)
	balanceTxt := string(data)
	balance, _ := strconv.ParseFloat(balanceTxt, 0644)
	return balance
}

// Bank app

func BankApp() {
	var accountBalance float64
	var userChoice int
	var amount float64
	fmt.Println("Welcome to bank app")
	for {
		fmt.Println("pick from the choices below to do something")
		fmt.Println("1. Make a deposit")
		fmt.Println("2. Make a Withdrawal")
		fmt.Println("3. Make a account balance check")
		fmt.Println("4. Exit the bank app")
		fmt.Scan(&userChoice)
		if userChoice == 1 {
			fmt.Println("Enter the amount you want to deposit")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += amount
			fmt.Println("Deposit Successful")
			fmt.Printf("Your new balance is $%.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		}else if  userChoice == 2 {
			fmt.Println("Enter the amount you want to withdraw")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}else if amount > accountBalance {
				fmt.Println("Insufficent Balance")
				fmt.Println("You can only withdraw less than or equal to: ", accountBalance)
				continue
			}
			accountBalance -= amount
			fmt.Println("Withdrawal Successful")
			fmt.Println("Your new balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		}else if userChoice == 3 {
			fmt.Println("Your current balance is: ", readBalanceFromFile())
		} else {
			fmt.Println("Exiting the bank app....Thank you for choosing us")
			return
		}
	}
}