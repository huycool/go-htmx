package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("hello world")

	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello World")
		// io.WriteString(w, r.Method)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}