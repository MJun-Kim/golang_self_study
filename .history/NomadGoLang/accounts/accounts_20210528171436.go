package accounts

import "errors"

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your amount
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}
