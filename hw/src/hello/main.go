package main

import (
	"encoding/json"
	"net/http"
	"poetry"
)

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	poemName := r.Form["name"][0]
	p, err := poetry.LoadPoem(poemName + ".txt")
	if err != nil {
		http.Error(w, "Poem not found", http.StatusNotFound)
	} else {
		enc := json.NewEncoder(w)
		enc.Encode(p)
	}
}

func main() {
	http.HandleFunc("/poem", poemHandler)
	http.ListenAndServe(":8088", nil)
}
