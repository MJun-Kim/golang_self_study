package main

import "net/http"

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	go pages := getPages()
}

func getPages() int {
	res, err := http.Get(baseUrl)
	return 0
}
