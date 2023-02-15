package main

type Account struct {
	balance int
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Depoist(amount int) {
	return a.
}