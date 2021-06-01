package main

import (
	"fmt"
	"std/github.com/kimminjun/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"First": "Firstword"}
	definition, err := dictionary.Search("Firt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
