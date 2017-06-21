package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"poetry"
	"sort"
	"strconv"
	"sync"
)

type protectedCache struct {
	sync.Mutex
	c map[string]poetry.Poem
}

var cache protectedCache

type config struct {
	Route       string
	BindAddress string   `json:"addr"`
	ValidPoems  []string `json:"valid"`
}

type poemWithTitle struct {
	Title     string
	Body      poetry.Poem
	WordCount string
	TheCount  int
}

var c config

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	poemName := r.Form["name"][0]

	p, ok := cache.c[poemName]
	if !ok {
		http.Error(w, "Poem not found (invalid)", http.StatusNotFound)
		return
	}

	sort.Sort(p[0])
	pwt := poemWithTitle{poemName, p,
		strconv.FormatInt(int64(p.NumWords()), 16), p.NumThe()}
	enc := json.NewEncoder(w)
	enc.Encode(pwt)
}

func main() {
	f, err := os.Open("config")
	if err != nil {
		fmt.Printf("Could not open JSON file.\nError: %s\n", err)
		os.Exit(1)
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	f.Close()
	if err != nil {
		fmt.Printf("Could not decode JSON file.\nError: %s\n", err)
		os.Exit(1)
	}

	cache.c = make(map[string]poetry.Poem)
	var wg sync.WaitGroup
	for _, name := range c.ValidPoems {
		wg.Add(1)
		go func(n string) {
			cache.Lock()
			cache.c[n], err = poetry.LoadPoem(n + ".txt")
			cache.Unlock()
			if err != nil {
				os.Exit(1)
			}
			wg.Done()
		}(name)
	}

	wg.Wait()
	http.HandleFunc(c.Route, poemHandler)
	http.ListenAndServe(c.BindAddress, nil)
}
