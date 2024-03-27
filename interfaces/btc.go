package main

import "errors"

type BitcoinAccount struct {
	ballance int
	fee      int
}

func NewBitcoin() *BitcoinAccount {
	return &BitcoinAccount{
		ballance: 0,
		fee:      30,
	}
}

func (b *BitcoinAccount) GetBallance() int {
	return b.ballance
}
func (b *BitcoinAccount) Deposit(amount int) {
	b.ballance += amount
}

func (b *BitcoinAccount) Withdraw(amount int) error {
	newBallance := b.ballance - (amount + b.fee)
	if newBallance < 0 {
		return errors.New("insufficient funds")
	}
	b.ballance = newBallance
	return nil
}
