package accounts

type Account struct {
	ID      int
	Balance float64
}

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}
