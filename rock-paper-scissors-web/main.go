package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// html := `<strong>Hello world</strong>`
	// w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, html)
	renderTemplate(w, "index.html")

}

func playRound(w http.ResponseWriter, r *http.Request) {
	round := rps.PlayRound(1)
	fmt.Println(round)

	out, err := json.MarshalIndent(round, "", "    ")

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(out)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)

	log.Println("starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		return
	}
}
