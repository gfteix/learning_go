package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// var ->  allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/*
In Go, when you call a function or a method the arguments are copied.

When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.

We can fix this with pointers. Pointers let us point to some values and then let us change them. So rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so that we can change the original values within it.
Note: no need to do '(*w).balance' as an explicit dereference, go automatically derefence for us
*/
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
