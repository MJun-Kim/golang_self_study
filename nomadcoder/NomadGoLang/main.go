package main

import (
	"basicgolang/NomadGoLang/accounts"
	"fmt"
)

// Whatever
func main() {
	account := accounts.NewAccount("MJun")
	account.Deposit(10)
	// account.ChangeOwner("haa")
	fmt.Println(account)

}
