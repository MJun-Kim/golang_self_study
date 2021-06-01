package main

import (
	"fmt"
	"std/github.com/kimminjun/NomadGoLang/banking"
)

// Whatever
func main() {
	account := banking.Account{Owner: "MJun"}
	account.Owner = "pape"
	fmt.Println(account)
}
