package main

import (
	"fmt"
	"std/github.com/kimminjun/NomadGoLang/accounts"
)

// Whatever
func main() {
	account := accounts.NewAccount("MJun")
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(20)
	fmt.Println(account.Balance())
}
