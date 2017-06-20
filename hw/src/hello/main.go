package main

import (
	"encoding/json"
	"net/http"
	"poetry"
)

type poemWithTitle struct {
	Title string
	Body  poetry.Poem
}

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	poemName := r.Form["name"][0]
	p, err := poetry.LoadPoem(poemName + ".txt")
	if err != nil {
		http.Error(w, "Poem not found", http.StatusNotFound)
	} else {
		pwt := poemWithTitle{poemName, p}
		enc := json.NewEncoder(w)
		enc.Encode(pwt)
	}
}

func main() {
	http.HandleFunc("/poem", poemHandler)
	http.ListenAndServe(":8088", nil)
}
