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
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())

}
