package main

import "errors"

type WellsFargo struct {
	ballance int
}

// constructor
func NewWellsFargo() *WellsFargo {
	return &WellsFargo{
		ballance: 0,
	}
}

func (w *WellsFargo) GetBallance() int {
	return w.ballance
}
func (w *WellsFargo) Deposit(amount int) {
	w.ballance += amount
}

func (w *WellsFargo) Withdraw(amount int) error {
	newBallance := w.ballance - amount
	if newBallance < 0 {
		return errors.New("insufficient funds")
	}
	w.ballance = newBallance
	return nil
}
