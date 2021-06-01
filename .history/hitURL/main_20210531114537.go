package main

import "fmt"

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagarm.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		fmt.Println(url)
	}
}
