package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"poetry"
	"sort"
	"strconv"
)

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

	found := false
	for _, v := range c.ValidPoems {
		if v == poemName {
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Poem not found (invalid)", http.StatusNotFound)
		return
	}

	p, err := poetry.LoadPoem(poemName + ".txt")
	if err != nil {
		http.Error(w, "Poem not found", http.StatusNotFound)
	} else {
		sort.Sort(p[0])
		pwt := poemWithTitle{poemName, p,
			strconv.FormatInt(int64(p.NumWords()), 16), p.NumThe()}
		enc := json.NewEncoder(w)
		enc.Encode(pwt)
	}
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

	http.HandleFunc(c.Route, poemHandler)
	http.ListenAndServe(c.BindAddress, nil)
}
