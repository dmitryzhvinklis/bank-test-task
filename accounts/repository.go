package accounts

import (
	"sync"
)

type AccountRepository struct {
	accounts map[int]BankAccount
	mutex    sync.Mutex
}

var repo *AccountRepository

func GetAccountRepository() *AccountRepository {
	if repo == nil {
		repo = &AccountRepository{
			accounts: make(map[int]BankAccount),
		}
	}
	return repo
}

func (r *AccountRepository) CreateAccount(id int) BankAccount {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	account := NewAccount(id)
	r.accounts[id] = account
	return account
}

func (r *AccountRepository) GetAccount(id int) (BankAccount, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	account, exists := r.accounts[id]
	return account, exists
}
