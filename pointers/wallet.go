package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Cannot withdraw. Insufficient funds!")

type Ethereum int

type Wallet struct {
	balance Ethereum
}

func (w *Wallet) Deposit(amount Ethereum) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Ethereum) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	} else {
		w.balance -= amount
		return nil
	}
}

func (w *Wallet) Balance() Ethereum {
	return w.balance
}

func (e Ethereum) String() string {
	return fmt.Sprintf("%d ETH", e)
}
