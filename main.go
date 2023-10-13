package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The idk", Director: "sdsdsds"},
				{Title: "skdsodks", Director: "sdsdsddsds"},
				{Title: "sdsdsd", Director: "dawasdsasad"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		//simulate latency
		time.Sleep(1 * time.Second)
		// To check wether htmx req was recieved
		log.Print("HTMX request was recieved")
		log.Print(r.Header.Get("HX-Request"))

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		//confirming the "submit" button
		fmt.Println("title added", title)
		fmt.Println("director added", director)

		// without using the shortcut thingy stuff A.K.A block "the action"
		// htmlStr := fmt.Sprintf(" <li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		// tmpl.Execute(w, nil)

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "The-action", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
