// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan string)
// 	people := [5]string{"nico", "flynn", "japanguy", "dal", "larry"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}

// 	for i := 0; i < len(people); i++ {
// 		fmt.Println("waiting for ", i)
// 		fmt.Println(<-c)
// 	}

// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- person + " is sexy"
// }

package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	// var results = make(map[string]string)
	c := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://academy.nomadcoders.co/",
		"https://www.instagram.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
