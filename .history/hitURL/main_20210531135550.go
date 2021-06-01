package main

import (
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	time.Sleep(time.Second * 10)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true

}

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var errRequestFailed = errors.New("Request failed")

// func main() {
// 	var results = make(map[string]string)

// 	urls := []string{
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://www.facebook.com/",
// 		"https://academy.nomadcoders.co/",
// 		"https://www.instagram.com/",
// 	}
// 	results["hello"] = "Hello"
// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}
// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}
// }

// func hitURL(url string) error {
// 	fmt.Println("Checking:", url)
// 	resp, err := http.Get(url)
// 	fmt.Println("response ====== ", resp.StatusCode, " ,,,   err ====== ", err)
// 	if err != nil || resp.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }
