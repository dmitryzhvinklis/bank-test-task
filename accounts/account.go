package accounts

import (
	"errors"
	"sync"
)

type account struct {
	ID      int
	balance float64
	mutex   sync.Mutex
}

func NewAccount(id int) BankAccount {
	return &account{
		ID: id,
	}
}

func (a *account) Deposit(amount float64) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	a.balance += amount
	return nil
}

func (a *account) Withdraw(amount float64) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *account) GetBalance() float64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.balance
}
