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
			{Title: "Aliens", Director: "Ridley Scott"},
			{Title: "Star Wars: New Hope", Director: "George Lucas"},
		},
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello World")
		// io.WriteString(w, r.Method)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, films)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Header.Get("HX-Request"))
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		fmt.Println(title + ":" + director)
		htmlStr := fmt.Sprintf(
			`<tr>
		<td>%s</td>
		<td>%s</td>
	  	</tr>`, title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
