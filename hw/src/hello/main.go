package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) isOk() bool {
	return w.err == nil
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()
	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func main() {
	w := &webPage{url: "http://www.oreilly.com/"}
	w.get()
	if w.isOk() {
		fmt.Printf("URL: %s Length: %d", w.url, len(w.body))
	} else {
		fmt.Printf("Something went wrong: %s", w.err)
	}
}
