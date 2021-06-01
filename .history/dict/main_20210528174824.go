package main

import (
	"fmt"
	"std/github.com/kimminjun/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "helloworld"
	fmt.Println(dictionary)
	dictionary.Delete("hello")
}
