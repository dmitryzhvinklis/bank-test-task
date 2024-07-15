package accounts

import (
	"errors"

	"github.com/dmitryzhvinklis/bank-test-task/utils"
)

func CreateAccount(id int) BankAccount {
	repo := GetAccountRepository()
	account := repo.CreateAccount(id)
	utils.LogOperation("CreateAccount", id)
	return account
}

func Deposit(id int, amount float64) error {
	repo := GetAccountRepository()
	account, exists := repo.GetAccount(id)
	if !exists {
		return errors.New("account not found")
	}
	err := account.Deposit(amount)
	utils.LogOperation("Deposit", id)
	return err
}

func Withdraw(id int, amount float64) error {
	repo := GetAccountRepository()
	account, exists := repo.GetAccount(id)
	if !exists {
		return errors.New("account not found")
	}
	err := account.Withdraw(amount)
	utils.LogOperation("Withdraw", id)
	return err
}

func GetBalance(id int) (float64, error) {
	repo := GetAccountRepository()
	account, exists := repo.GetAccount(id)
	if !exists {
		return 0, errors.New("account not found")
	}
	balance := account.GetBalance()
	utils.LogOperation("GetBalance", id)
	return balance, nil
}
