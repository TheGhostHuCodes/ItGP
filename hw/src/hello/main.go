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

func generator(url string, urlCh chan string) {
	urlCh <- url
}

func worker(urlCh chan string, responseCh chan string, id int) {
	for {
		url := <-urlCh
		length, err := getPage(url)
		if err == nil {
			responseCh <- fmt.Sprintf("(worker:%d)%s has length %d", id, url, length)
		} else {
			responseCh <- fmt.Sprintf("(worker:%d)Error getting %s: %s", id, url, err)
		}
	}
}

func main() {
	urls := []string{"http://www.google.com/", "http://www.yahoo.com",
		"http://www.bing.com", "http://bbc.co.uk", "http://www.oreilly.com/"}
	urlCh := make(chan string)
	responseCh := make(chan string)
	for i := 0; i < 10; i++ {
		go worker(urlCh, responseCh, i)
	}

	for _, url := range urls {
		go generator(url, urlCh)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-responseCh)
	}
}
