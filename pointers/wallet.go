package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
} 

type Wallet struct {
  balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
  w.balance += amount
}

var ErrInsufficientAmount = errors.New("Insufficient amount")

func (w *Wallet) Withdraw(amount Bitcoin) error {
  if amount > w.Balance() {
    return ErrInsufficientAmount 
  }
  w.balance -= amount
  return nil
}

func (w *Wallet) Balance() Bitcoin {
  return w.balance
}
