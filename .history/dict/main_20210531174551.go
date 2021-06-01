package main

import (
	"basicgolang/dict/mydict"
	"fmt"
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
