package main

import (
	"fmt"
	"log"
	"std/github.com/kimminjun/NomadGoLang/accounts"
)

// Whatever
func main() {
	account := accounts.NewAccount("MJun")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())

}
