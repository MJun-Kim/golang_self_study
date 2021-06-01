package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://academy.nomadcoders.co/",
		"https://www.instagarm.com/",
	}
	results["hello"] = "Hello"
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	fmt.Println(results)
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	fmt.Println("response ====== ", resp.StatusCode, " ,,,   err ====== ", err)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
