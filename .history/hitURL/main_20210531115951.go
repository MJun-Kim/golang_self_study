package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

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
		hitURL(url)
	}
}

func hitURL(url string) error {
	resp, err := http.Get(url)
	fmt.Println("response ====== ", resp, " ,,,   err ====== ", err)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
