package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	return len(body), nil
}

func getter(url string, response chan string) {
	length, err := getPage(url)
	if err == nil {
		response <- fmt.Sprintf("%s has length %d", url, length)
	}
}

func main() {
	urls := []string{"http://www.google.com/", "http://www.yahoo.com",
		"http://www.bing.com", "http://bbc.co.uk"}
	response := make(chan string)
	for _, url := range urls {
		go getter(url, response)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-response)
	}
}
