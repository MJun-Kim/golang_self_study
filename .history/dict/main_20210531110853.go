package main

import (
	"fmt"
	"std/github.com/kimminjun/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	} else {

	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

}
