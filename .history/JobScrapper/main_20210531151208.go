package main

import "net/http"

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := petPages()
}

func petPages() int {
	res, err := http.Get
	return 0
}
