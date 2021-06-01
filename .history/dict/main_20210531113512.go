package main

import (
	"fmt"
	"std/github.com/kimminjun/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

}
