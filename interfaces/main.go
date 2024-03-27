package main

import "fmt"

type BackAccoutI interface {
	GetBallance() int //100 = 1dolar
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	myAccounts := []BackAccoutI{
		NewWellsFargo(),
		NewBitcoin(),
	}

	for _, account := range myAccounts {
		account.Deposit(300)
		if err := account.Withdraw(70); err != nil {
			fmt.Println("not enough money")
		}
		ballance := account.GetBallance()
		fmt.Println("ballance :", ballance)
	}
}
