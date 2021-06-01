// A concurrent prime sieve

package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 1000; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

// package main

// import (
// 	"fmt"
// 	"std/github.com/kimminjun/dict/mydict"
// )

// func main() {
// 	dictionary := mydict.Dictionary{"First": "Firstword"}
// 	definition, err := dictionary.Search("First")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(definition)
// 	}
// }
